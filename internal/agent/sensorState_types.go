// Code generated by "stringer -type=deviceClassType,stateClassType -output sensorState_types.go -trimprefix deviceClass"; DO NOT EDIT.

package agent

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[apparent_power-1]
	_ = x[aqi-2]
	_ = x[atmospheric_pressure-3]
	_ = x[deviceClassBattery-4]
	_ = x[carbon_dioxide-5]
	_ = x[carbon_monoxide-6]
	_ = x[current-7]
	_ = x[data_rate-8]
	_ = x[data_size-9]
	_ = x[date-10]
	_ = x[distance-11]
	_ = x[duration-12]
	_ = x[energy-13]
	_ = x[enum-14]
	_ = x[frequency-15]
	_ = x[gas-16]
	_ = x[humidity-17]
	_ = x[illuminance-18]
	_ = x[irradiance-19]
	_ = x[moisture-20]
	_ = x[monetary-21]
	_ = x[nitrogen_dioxide-22]
	_ = x[nitrogen_monoxide-23]
	_ = x[nitrous_oxide-24]
	_ = x[ozone-25]
	_ = x[pm1-26]
	_ = x[pm25-27]
	_ = x[pm10-28]
	_ = x[power_factor-29]
	_ = x[deviceClassPower-30]
	_ = x[precipitation-31]
	_ = x[precipitation_intensity-32]
	_ = x[pressure-33]
	_ = x[reactive_power-34]
	_ = x[signal_strength-35]
	_ = x[sound_pressure-36]
	_ = x[speed-37]
	_ = x[sulphur_dioxide-38]
	_ = x[deviceClassTemperature-39]
	_ = x[timestamp-40]
	_ = x[volatile_organic_compounds-41]
	_ = x[voltage-42]
	_ = x[volume-43]
	_ = x[water-44]
	_ = x[weight-45]
	_ = x[wind_speed-46]
}

const _deviceClassType_name = "apparent_poweraqiatmospheric_pressureBatterycarbon_dioxidecarbon_monoxidecurrentdata_ratedata_sizedatedistancedurationenergyenumfrequencygashumidityilluminanceirradiancemoisturemonetarynitrogen_dioxidenitrogen_monoxidenitrous_oxideozonepm1pm25pm10power_factorPowerprecipitationprecipitation_intensitypressurereactive_powersignal_strengthsound_pressurespeedsulphur_dioxideTemperaturetimestampvolatile_organic_compoundsvoltagevolumewaterweightwind_speed"

var _deviceClassType_index = [...]uint16{0, 14, 17, 37, 44, 58, 73, 80, 89, 98, 102, 110, 118, 124, 128, 137, 140, 148, 159, 169, 177, 185, 201, 218, 231, 236, 239, 243, 247, 259, 264, 277, 300, 308, 322, 337, 351, 356, 371, 382, 391, 417, 424, 430, 435, 441, 451}

func (i deviceClassType) String() string {
	i -= 1
	if i < 0 || i >= deviceClassType(len(_deviceClassType_index)-1) {
		return "deviceClassType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _deviceClassType_name[_deviceClassType_index[i]:_deviceClassType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[measurement-46]
	_ = x[total-47]
	_ = x[total_increasing-48]
}

const _stateClassType_name = "measurementtotaltotal_increasing"

var _stateClassType_index = [...]uint8{0, 11, 16, 32}

func (i stateClassType) String() string {
	i -= 46
	if i < 0 || i >= stateClassType(len(_stateClassType_index)-1) {
		return "stateClassType(" + strconv.FormatInt(int64(i+46), 10) + ")"
	}
	return _stateClassType_name[_stateClassType_index[i]:_stateClassType_index[i+1]]
}
