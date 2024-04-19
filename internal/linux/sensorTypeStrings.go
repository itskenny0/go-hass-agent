// Code generated by "stringer -type=SensorTypeValue -output sensorTypeStrings.go -linecomment"; DO NOT EDIT.

package linux

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SensorAppActive-1]
	_ = x[SensorAppRunning-2]
	_ = x[SensorBattType-3]
	_ = x[SensorBattPercentage-4]
	_ = x[SensorBattTemp-5]
	_ = x[SensorBattVoltage-6]
	_ = x[SensorBattEnergy-7]
	_ = x[SensorBattEnergyRate-8]
	_ = x[SensorBattState-9]
	_ = x[SensorBattNativePath-10]
	_ = x[SensorBattLevel-11]
	_ = x[SensorBattModel-12]
	_ = x[SensorMemTotal-13]
	_ = x[SensorMemAvail-14]
	_ = x[SensorMemUsed-15]
	_ = x[SensorMemPc-16]
	_ = x[SensorSwapTotal-17]
	_ = x[SensorSwapUsed-18]
	_ = x[SensorSwapFree-19]
	_ = x[SensorSwapPc-20]
	_ = x[SensorConnectionState-21]
	_ = x[SensorConnectionID-22]
	_ = x[SensorConnectionDevices-23]
	_ = x[SensorConnectionType-24]
	_ = x[SensorConnectionIPv4-25]
	_ = x[SensorConnectionIPv6-26]
	_ = x[SensorAddressIPv4-27]
	_ = x[SensorAddressIPv6-28]
	_ = x[SensorWifiSSID-29]
	_ = x[SensorWifiFrequency-30]
	_ = x[SensorWifiSpeed-31]
	_ = x[SensorWifiStrength-32]
	_ = x[SensorWifiHWAddress-33]
	_ = x[SensorBytesSent-34]
	_ = x[SensorBytesRecv-35]
	_ = x[SensorBytesSentRate-36]
	_ = x[SensorBytesRecvRate-37]
	_ = x[SensorPowerProfile-38]
	_ = x[SensorBoottime-39]
	_ = x[SensorUptime-40]
	_ = x[SensorLoad1-41]
	_ = x[SensorLoad5-42]
	_ = x[SensorLoad15-43]
	_ = x[SensorCPUPc-44]
	_ = x[SensorScreenLock-45]
	_ = x[SensorLaptopLid-46]
	_ = x[SensorProblem-47]
	_ = x[SensorKernel-48]
	_ = x[SensorDistribution-49]
	_ = x[SensorVersion-50]
	_ = x[SensorUsers-51]
	_ = x[SensorDeviceTemp-52]
	_ = x[SensorPowerState-53]
	_ = x[SensorAccentColor-54]
	_ = x[SensorColorScheme-55]
	_ = x[SensorDiskReads-56]
	_ = x[SensorDiskWrites-57]
	_ = x[SensorDiskReadRate-58]
	_ = x[SensorDiskWriteRate-59]
}

const _SensorTypeValue_name = "Active AppRunning AppsBattery TypeBattery LevelBattery TemperatureBattery VoltageBattery EnergyBattery PowerBattery StateBattery PathBattery LevelBattery ModelMemory TotalMemory AvailableMemory UsedMemory UsageSwap Memory TotalSwap Memory UsedSwap Memory FreeSwap UsageConnection StateConnection IDConnection DeviceConnection TypeConnection IPv4Connection IPv6IPv4 AddressIPv6 AddressWi-Fi SSIDWi-Fi FrequencyWi-Fi Link SpeedWi-Fi Signal StrengthWi-Fi BSSIDBytes SentBytes ReceivedBytes Sent ThroughputBytes Received ThroughputPower ProfileLast RebootUptimeCPU load average (1 min)CPU load average (5 min)CPU load average (15 min)CPU UsageScreen LockLaptop LidProblemsKernel VersionDistribution NameDistribution VersionCurrent UsersTemperaturePower StateAccent ColorColor Scheme TypeDisk ReadsDisk WritesDisk Read RateDisk Write Rate"

var _SensorTypeValue_index = [...]uint16{0, 10, 22, 34, 47, 66, 81, 95, 108, 121, 133, 146, 159, 171, 187, 198, 210, 227, 243, 259, 269, 285, 298, 315, 330, 345, 360, 372, 384, 394, 409, 425, 446, 457, 467, 481, 502, 527, 540, 551, 557, 581, 605, 630, 639, 650, 660, 668, 682, 699, 719, 732, 743, 754, 766, 783, 793, 804, 818, 833}

func (i SensorTypeValue) String() string {
	i -= 1
	if i < 0 || i >= SensorTypeValue(len(_SensorTypeValue_index)-1) {
		return "SensorTypeValue(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _SensorTypeValue_name[_SensorTypeValue_index[i]:_SensorTypeValue_index[i+1]]
}
