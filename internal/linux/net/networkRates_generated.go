// Code generated by "stringer -type=rateSensor -output networkRates_generated.go -linecomment"; DO NOT EDIT.

package net

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[bytesSent-4]
	_ = x[bytesRecv-5]
	_ = x[bytesSentRate-6]
	_ = x[bytesRecvRate-7]
}

const _rateSensor_name = "Bytes SentBytes ReceivedBytes Sent ThroughputBytes Received Throughput"

var _rateSensor_index = [...]uint8{0, 10, 24, 45, 70}

func (i rateSensor) String() string {
	i -= 4
	if i < 0 || i >= rateSensor(len(_rateSensor_index)-1) {
		return "rateSensor(" + strconv.FormatInt(int64(i+4), 10) + ")"
	}
	return _rateSensor_name[_rateSensor_index[i]:_rateSensor_index[i+1]]
}