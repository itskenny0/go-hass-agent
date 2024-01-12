// Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package api

import (
	"context"
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/joshuar/go-hass-agent/internal/agent/config"
	"github.com/lxzan/gws"
	"github.com/rs/zerolog/log"
)

const PingInterval = time.Minute

type websocketMsg struct {
	Type           string `json:"type"`
	WebHookID      string `json:"webhook_id,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	ID             uint64 `json:"id,omitempty"`
	SupportConfirm bool   `json:"support_confirm,omitempty"`
}

func (m *websocketMsg) send(conn *gws.Conn) error {
	msg, err := json.Marshal(&m)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(gws.OpcodeText, msg)
	if err != nil {
		return err
	}
	return nil
}

type websocketResponse struct {
	Result       any                   `json:"result,omitempty"`
	Error        ResponseError         `json:"error,omitempty"`
	Type         string                `json:"type"`
	HAVersion    string                `json:"ha_version,omitempty"`
	Notification websocketNotification `json:"event,omitempty"`
	ID           uint64                `json:"id,omitempty"`
	Success      bool                  `json:"success,omitempty"`
}

type websocketNotification struct {
	Data      any      `json:"data,omitempty"`
	Message   string   `json:"message"`
	Title     string   `json:"title,omitempty"`
	ConfirmID string   `json:"confirm_id,omitempty"`
	Target    []string `json:"target,omitempty"`
}

func StartWebsocket(ctx context.Context, notifyCh chan [2]string) {
	cfg := config.FetchFromContext(ctx)
	if cfg == nil {
		log.Warn().Msg("Cannot retrieve config. Cannot listen for notifications.")
		return
	}

	var websocketURL string
	if err := cfg.Get(config.PrefWebsocketURL, &websocketURL); err != nil {
		log.Warn().Err(err).Msg("Could not retrieve websocket URL from config.")
		return
	}
	var socket *gws.Conn
	var err error

	retryFunc := func() error {
		var resp *http.Response
		socket, resp, err = gws.NewClient(
			newWebsocket(cfg, notifyCh),
			&gws.ClientOption{Addr: websocketURL})
		if err != nil {
			log.Error().Err(err).
				Msg("Could not connect to websocket.")
			return err
		}
		defer resp.Body.Close()
		return nil
	}
	err = backoff.Retry(retryFunc, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		log.Error().Err(err).
			Msg("Could not connect to websocket.")
		return
	}
	log.Trace().Caller().Msg("Websocket connection established.")

	go func() {
		<-ctx.Done()
		socket.WriteClose(1000, nil)
	}()
	socket.ReadLoop()
}

type WebSocket struct {
	notifyCh  chan [2]string
	doneCh    chan struct{}
	token     string
	webhookID string
	nextID    uint64
}

func newWebsocket(cfg config.Config, notifyCh chan [2]string) *WebSocket {
	var token, webhookID string
	if err := cfg.Get(config.PrefToken, &token); err != nil {
		log.Warn().Err(err).Msg("Could not retrieve token from config.")
		return nil
	}
	if err := cfg.Get(config.PrefWebhookID, &webhookID); err != nil {
		log.Warn().Err(err).Msg("Could not retrieve webhookID from config.")
		return nil
	}

	ws := &WebSocket{
		notifyCh:  notifyCh,
		doneCh:    make(chan struct{}),
		token:     token,
		webhookID: webhookID,
	}
	return ws
}

func (c *WebSocket) newAuthMsg() *websocketMsg {
	return &websocketMsg{
		Type:        "auth",
		AccessToken: c.token,
	}
}

func (c *WebSocket) newRegistrationMsg() *websocketMsg {
	return &websocketMsg{
		Type:           "mobile_app/push_notification_channel",
		ID:             atomic.LoadUint64(&c.nextID),
		WebHookID:      c.webhookID,
		SupportConfirm: false,
	}
}

func (c *WebSocket) newPingMsg() *websocketMsg {
	return &websocketMsg{
		Type: "ping",
		ID:   atomic.LoadUint64(&c.nextID),
	}
}

func (c *WebSocket) OnError(_ *gws.Conn, err error) {
	log.Error().Err(err).
		Msg("Error on websocket")
}

func (c *WebSocket) OnClose(_ *gws.Conn, err error) {
	if err.Error() != "" {
		log.Error().Err(err).Msg("Websocket connection closed with error.")
	}
	close(c.doneCh)
}

func (c *WebSocket) OnPong(_ *gws.Conn, payload []byte) {
	log.Trace().Caller().Msg("Received pong on websocket")
}

func (c *WebSocket) OnOpen(socket *gws.Conn) {
	log.Trace().Caller().Msg("Websocket opened.")
	go c.keepAlive(socket)
}

func (c *WebSocket) OnPing(socket *gws.Conn, payload []byte) {
	log.Trace().Msg("Received ping on websocket.")
}

func (c *WebSocket) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	response := &websocketResponse{
		Success: true,
	}
	if err := json.Unmarshal(message.Bytes(), &response); err != nil {
		log.Error().Err(err).
			Msgf("Failed to unmarshall response %s.", message.Data.String())
		return
	}
	atomic.AddUint64(&c.nextID, 1)
	var r *websocketMsg
	switch response.Type {
	case "event":
		c.notifyCh <- [2]string{response.Notification.Title, response.Notification.Message}
	case "result":
		if !response.Success {
			log.Error().
				Msgf("Received error on websocket, %s: %s.", response.Error.ErrorCode, response.Error.ErrorMsg)
			if response.Error.ErrorCode == "id_reuse" {
				log.Warn().
					Msg("id_reuse error, attempting manual increment.")
				atomic.AddUint64(&c.nextID, 1)
			}
		}
	case "auth_required":
		log.Trace().Caller().
			Msg("Requesting authorisation for websocket.")
		r = c.newAuthMsg()
	case "auth_ok":
		log.Trace().Caller().
			Msg("Registering app for push notifications.")
		r = c.newRegistrationMsg()
	case "pong":
		b, err := json.Marshal(response)
		if err != nil {
			log.Error().Err(err).Msg("Unable to unmarshal.")
		}
		c.OnPong(socket, b)
	default:
		log.Warn().Caller().Msgf("Unhandled websocket response %v.", response.Type)
	}
	if r != nil {
		err := r.send(socket)
		if err != nil {
			log.Error().Err(err).
				Msg("Unable to send websocket message.")
		}
	}
}

func (c *WebSocket) keepAlive(conn *gws.Conn) {
	ticker := time.NewTicker(PingInterval)
	for {
		select {
		case <-c.doneCh:
			return
		case <-ticker.C:
			log.Trace().Caller().
				Msg("Sending ping on websocket")
			if err := conn.SetDeadline(time.Now().Add(2 * PingInterval)); err != nil {
				log.Error().Err(err).
					Msg("Error setting deadline on websocket.")
				return
			}
			msg := c.newPingMsg()
			if err := msg.send(conn); err != nil {
				log.Error().Err(err).
					Msg("Error sending ping.")
			}
		}
	}
}
