package agent

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/joshuar/go-hass-agent/internal/device"
	"github.com/joshuar/go-hass-agent/internal/hass"
	"github.com/rs/zerolog/log"
)

// batterySensor is a specific type of sensorState for battery sensors
type batterySensor sensorState

func newBatterySensor(batteryID string, sensorType string) *batterySensor {
	var sensorName, sensorID string
	var stateClass stateClassType
	var deviceClass deviceClassType
	switch sensorType {
	case "Percentage":
		sensorName = batteryID + " battery level"
		sensorID = batteryID + "_battery_level"
		stateClass = measurement
		deviceClass = deviceClassBattery
	case "Temperature":
		sensorName = batteryID + " battery temperature"
		sensorID = batteryID + "_battery_temperature"
		stateClass = measurement
		deviceClass = deviceClassTemperature
	case "EnergyRate":
		sensorName = batteryID + " battery power"
		sensorID = batteryID + "_battery_power"
		stateClass = measurement
		deviceClass = deviceClassPower
	case "BatteryStatus":
		fallthrough
	case "BatteryLevel":
		sensorName = batteryID + " " + strcase.ToDelimited(sensorType, ' ')
		sensorID = batteryID + "_" + strings.ToLower(strcase.ToSnake(sensorType))
		stateClass = 0
		deviceClass = 0
	default:
		sensorName = batteryID + " battery " + strcase.ToDelimited(sensorType, ' ')
		sensorID = batteryID + "_" + strings.ToLower(strcase.ToSnake(sensorType))
		stateClass = 0
		deviceClass = 0
	}
	return &batterySensor{
		name:        sensorName,
		entityID:    sensorID,
		deviceClass: deviceClass,
		stateClass:  stateClass,
	}
}

// Ensure a batterySensor satisfies the sensor interface so it can be treated as
// a sensor

func (b *batterySensor) Attributes() interface{} {
	return b.attributes
}

func (b *batterySensor) DeviceClass() string {
	if b.deviceClass != 0 {
		return b.deviceClass.String()
	} else {
		return ""
	}
}

func (b *batterySensor) Icon() string {
	switch b.deviceClass {
	case deviceClassBattery:
		if b.state.(float64) >= 95 {
			return "mdi:battery"
		} else {
			return fmt.Sprintf("mdi:battery-%d", int(math.Round(b.state.(float64)/10)*10))
		}
	case deviceClassPower:
		if math.Signbit(b.state.(float64)) {
			return "mdi:battery-minus"
		} else {
			return "mdi:battery-plus"
		}
	default:
		return "mdi:battery"
	}
}

func (b *batterySensor) Name() string {
	return b.name
}

func (b *batterySensor) State() interface{} {
	switch b.deviceClass {
	case deviceClassBattery:
		return b.state.(float64)
	default:
		return b.state
	}
}

func (b *batterySensor) SensorType() string {
	return "sensor"
}

func (b *batterySensor) UniqueID() string {
	return b.entityID
}

func (b *batterySensor) UnitOfMeasurement() string {
	switch b.deviceClass {
	case deviceClassBattery:
		return "%"
	case deviceClassTemperature:
		return "°C"
	case deviceClassPower:
		return "W"
	default:
		return ""
	}
}

func (b *batterySensor) StateClass() string {
	if b.stateClass != 0 {
		return b.stateClass.String()
	} else {
		return ""
	}
}

func (b *batterySensor) EntityCategory() string {
	return "diagnostic"
}

func (b *batterySensor) Disabled() bool {
	return b.disabled
}

func (b *batterySensor) Registered() bool {
	return b.registered
}

// Ensure that a batterySensor satisfies the hass.Request interface so its data
// can be sent as a request to HA

func (b *batterySensor) RequestType() hass.RequestType {
	if b.Registered() {
		return hass.RequestTypeUpdateSensorStates
	}
	return hass.RequestTypeRegisterSensor
}

func (b *batterySensor) RequestData() interface{} {
	return hass.MarshallSensorData(b)
}

func (b *batterySensor) ResponseHandler(rawResponse interface{}) {
	if rawResponse == nil {
		log.Debug().Caller().Msg("No response data.")
	} else {
		response := rawResponse.(map[string]interface{})
		if v, ok := response["success"]; ok {
			if v.(bool) && !b.registered {
				b.registered = true
				log.Debug().Caller().Msgf("Sensor %s registered.", b.Name())
			}
		}
		if v, ok := response[b.entityID]; ok {
			status := v.(map[string]interface{})
			if !status["success"].(bool) {
				error := status["error"].(map[string]interface{})
				log.Error().Msgf("Could not update sensor %s, %s: %s", b.Name(), error["code"], error["message"])
			} else {
				log.Debug().Msgf("Sensor %s updated. State is now: %v", b.Name(), b.State())
			}
			if v, ok := status["is_disabled"]; ok {
				switch v.(bool) {
				case true:
					log.Debug().Msgf("Sensor %s has been disabled.", b.Name())
					b.disabled = true
				case false:
					log.Debug().Msgf("Sensor %s has been enabled.", b.Name())
					b.disabled = false
				}
			}
		}
	}
}

func (agent *Agent) runBatterySensorWorker(ctx context.Context) {

	updateCh := make(chan interface{})
	defer close(updateCh)

	sensors := make(map[string]*batterySensor)

	go device.BatteryUpdater(ctx, updateCh)

	for {
		select {
		case data := <-updateCh:
			update := data.(sensorUpdate)
			sensorID := update.Device() + update.Type()
			if _, ok := sensors[sensorID]; !ok {
				sensors[sensorID] = newBatterySensor(update.Device(), update.Type())
			}
			sensors[sensorID].state = update.Value()
			sensors[sensorID].attributes = update.ExtraValues()
			go hass.APIRequest(ctx, sensors[sensorID])
		case <-ctx.Done():
			log.Debug().Caller().
				Msg("Cleaning up battery sensors.")
			return
		}
	}
}
