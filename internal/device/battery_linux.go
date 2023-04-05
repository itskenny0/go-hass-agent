package device

import (
	"context"

	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
)

//go:generate stringer -type=BatteryProp -output battery_props_linux.go -trimprefix batt

const (
	upowerDBusDest         = "org.freedesktop.UPower"
	upowerDBusPath         = "/org/freedesktop/UPower"
	upowerGetDevicesMethod = "org.freedesktop.UPower.EnumerateDevices"

	battType BatteryProp = iota
	Percentage
	Temperature
	Voltage
	Energy
	EnergyRate
	battState
	NativePath
	BatteryLevel
)

type BatteryProp int

type upowerBattery struct {
	dBusPath dbus.ObjectPath
	props    map[BatteryProp]dbus.Variant
}

func (b *upowerBattery) updateProp(api *deviceAPI, prop BatteryProp) {
	propValue, err := api.GetDBusProp(systemBus, upowerDBusDest, b.dBusPath, "org.freedesktop.UPower.Device."+prop.String())
	if err != nil {
		log.Debug().Caller().Msgf(err.Error())
	}
	b.props[prop] = propValue
}

func (b *upowerBattery) getProp(prop BatteryProp) interface{} {
	return b.props[prop].Value()
}

func (b *upowerBattery) marshallStateUpdate(api *deviceAPI, prop BatteryProp) *upowerBatteryState {
	log.Debug().Caller().Msgf("Marshalling update for %v for battery %v", prop.String(), b.getProp(NativePath).(string))
	state := &upowerBatteryState{
		batteryID: b.getProp(NativePath).(string),
		prop: upowerBatteryProp{
			kind:  prop,
			value: b.getProp(prop),
		},
	}
	switch prop {
	case EnergyRate:
		b.updateProp(api, Voltage)
		b.updateProp(api, Energy)
		state.attributes = &struct {
			Voltage float64 `json:"Voltage"`
			Energy  float64 `json:"Energy"`
		}{
			Voltage: b.getProp(Voltage).(float64),
			Energy:  b.getProp(Energy).(float64),
		}
	case Percentage:
		fallthrough
	case BatteryLevel:
		state.attributes = &struct {
			Type string `json:"Battery Type"`
		}{
			Type: stringType(b.getProp(battType).(uint32)),
		}
	}
	return state
}

type upowerBatteryProp struct {
	kind  BatteryProp
	value interface{}
}

type upowerBatteryState struct {
	batteryID  string
	prop       upowerBatteryProp
	attributes interface{}
}

func (state *upowerBatteryState) ID() string {
	return state.batteryID
}

func (state *upowerBatteryState) Type() BatteryProp {
	return state.prop.kind
}

func (state *upowerBatteryState) Value() interface{} {
	switch state.prop.kind {
	case Voltage:
		fallthrough
	case Temperature:
		fallthrough
	case Energy:
		fallthrough
	case EnergyRate:
		fallthrough
	case Percentage:
		return state.prop.value.(float64)
	case battState:
		return stringState(state.prop.value.(uint32))
	case BatteryLevel:
		return stringLevel(state.prop.value.(uint32))
	default:
		return state.prop.value.(string)
	}
}

func (state *upowerBatteryState) ExtraValues() interface{} {
	return state.attributes
}

func stringState(state uint32) string {
	switch state {
	case 1:
		return "Charging"
	case 2:
		return "Discharging"
	case 3:
		return "Empty"
	case 4:
		return "Fully Charged"
	case 5:
		return "Pending Charge"
	case 6:
		return "Pending Discharge"
	default:
		return "Unknown"
	}
}

func stringType(t uint32) string {
	switch t {
	case 0:
		return "Unknown"
	case 1:
		return "Line Power"
	case 2:
		return "Battery"
	case 3:
		return "Ups"
	case 4:
		return "Monitor"
	case 5:
		return "Mouse"
	case 6:
		return "Keyboard"
	case 7:
		return "Pda"
	case 8:
		return "Phone"
	default:
		return "Unknown"
	}
}

func stringLevel(l uint32) string {
	switch l {
	case 0:
		return "Unknown"
	case 1:
		return "None"
	case 3:
		return "Low"
	case 4:
		return "Critical"
	case 6:
		return "Normal"
	case 7:
		return "High"
	case 8:
		return "Full"
	default:
		return "Unknown"
	}
}

func BatteryUpdater(ctx context.Context, status chan interface{}) {

	deviceAPI, deviceAPIExists := FromContext(ctx)
	if !deviceAPIExists {
		log.Debug().Caller().
			Msg("Could not connect to DBus to monitor batteries.")
		return
	}

	batteryList, err := deviceAPI.GetDBusData(systemBus, upowerDBusDest, upowerDBusPath, upowerGetDevicesMethod)
	if err != nil {
		log.Debug().Caller().
			Msgf("Unable to find all battery devices: %v", err)
		return
	}

	batteryTracker := make(map[string]*upowerBattery)
	for _, v := range batteryList.([]dbus.ObjectPath) {

		// Track this battery in batteryTracker.
		batteryID := string(v)
		batteryTracker[batteryID] = &upowerBattery{
			dBusPath: v,
		}
		batteryTracker[batteryID].props = make(map[BatteryProp]dbus.Variant)
		batteryTracker[batteryID].updateProp(deviceAPI, NativePath)
		batteryTracker[batteryID].updateProp(deviceAPI, battType)

		// Standard battery properties as sensors
		for _, prop := range []BatteryProp{battState} {
			batteryTracker[batteryID].updateProp(deviceAPI, prop)
			stateUpdate := batteryTracker[batteryID].marshallStateUpdate(deviceAPI, prop)
			if stateUpdate != nil {
				status <- stateUpdate
			}
		}

		// For some battery types, track additional properties as sensors
		if batteryTracker[batteryID].getProp(battType).(uint32) == 2 {
			for _, prop := range []BatteryProp{Percentage, Temperature, EnergyRate} {
				batteryTracker[batteryID].updateProp(deviceAPI, prop)
				stateUpdate := batteryTracker[batteryID].marshallStateUpdate(deviceAPI, prop)
				if stateUpdate != nil {
					status <- stateUpdate
				}
			}
		} else {
			batteryTracker[batteryID].updateProp(deviceAPI, BatteryLevel)
			stateUpdate := batteryTracker[batteryID].marshallStateUpdate(deviceAPI, BatteryLevel)
			if stateUpdate != nil {
				status <- stateUpdate
			}
		}

		// Create a DBus signal match to watch for property changes for this
		// battery. If a property changes, check it is one we want to track and
		// if so, update the battery's state in batteryTracker and send the
		// update back to Home Assistant.
		batteryChangeSignal := &DBusWatchRequest{
			bus: systemBus,
			match: DBusSignalMatch{
				path: v,
				intr: "org.freedesktop.DBus.Properties",
			},
			event: "org.freedesktop.DBus.Properties.PropertiesChanged",
			eventHandler: func(s *dbus.Signal) {
				log.Debug().Caller().Msg("Recieved changed battery state.")
				batteryID := string(s.Path)
				props := s.Body[1].(map[string]dbus.Variant)
				for propName, propValue := range props {
					for BatteryProp := range batteryTracker[batteryID].props {
						if propName == BatteryProp.String() {
							batteryTracker[batteryID].props[BatteryProp] = propValue
							log.Debug().Caller().
								Msgf("Updating battery property %v to %v", BatteryProp.String(), propValue.Value())
							stateUpdate := batteryTracker[batteryID].marshallStateUpdate(deviceAPI, BatteryProp)
							if stateUpdate != nil {
								status <- stateUpdate
							}
						}
					}
				}
			},
		}
		deviceAPI.WatchEvents <- batteryChangeSignal
	}

	<-status
	log.Debug().Caller().
		Msg("Stopping Linux battery updater.")
}
