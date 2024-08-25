// Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

//revive:disable:unused-receiver
package cpu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joshuar/go-hass-agent/internal/hass/sensor"
	"github.com/joshuar/go-hass-agent/internal/hass/sensor/types"
	"github.com/joshuar/go-hass-agent/internal/linux"
	"github.com/joshuar/go-hass-agent/pkg/linux/dbusx"
)

const (
	loadAvgIcon = "mdi:chip"
	loadAvgUnit = "load"

	loadAvgUpdateInterval = time.Minute
	loadAvgUpdateJitter   = 5 * time.Second

	loadAvgsTotal = 3

	loadAvgsWorkerID = "load_averages_sensors"
)

var ErrParseLoadAvgs = errors.New("error parsing load averages")

type loadavgSensor struct {
	linux.Sensor
}

type loadAvgsSensorWorker struct {
	path     string
	loadAvgs []*loadavgSensor
}

func (w *loadAvgsSensorWorker) Interval() time.Duration { return loadAvgUpdateInterval }

func (w *loadAvgsSensorWorker) Jitter() time.Duration { return loadAvgUpdateJitter }

func (w *loadAvgsSensorWorker) Sensors(_ context.Context, _ time.Duration) ([]sensor.Details, error) {
	sensors := make([]sensor.Details, loadAvgsTotal)

	loadAvgData, err := os.ReadFile(w.path)
	if err != nil {
		return nil, fmt.Errorf("fetch load averages: %w", err)
	}

	loadAvgs, err := parseLoadAvgs(loadAvgData)
	if err != nil {
		return nil, fmt.Errorf("parse load averages: %w", err)
	}

	for idx := range loadAvgs {
		w.loadAvgs[idx].Value = loadAvgs[idx]
		sensors[idx] = w.loadAvgs[idx]
	}

	return sensors, nil
}

func newLoadAvgSensors() []*loadavgSensor {
	sensors := make([]*loadavgSensor, loadAvgsTotal)

	for idx, loadType := range []linux.SensorTypeValue{linux.SensorLoad1, linux.SensorLoad5, linux.SensorLoad15} {
		loadAvgSensor := &loadavgSensor{}
		loadAvgSensor.IconString = loadAvgIcon
		loadAvgSensor.UnitsString = loadAvgUnit
		loadAvgSensor.SensorSrc = linux.DataSrcProcfs
		loadAvgSensor.StateClassValue = types.StateClassMeasurement

		switch loadType { //nolint:exhaustive
		case linux.SensorLoad1:
			loadAvgSensor.SensorTypeValue = linux.SensorLoad1
		case linux.SensorLoad5:
			loadAvgSensor.SensorTypeValue = linux.SensorLoad5
		case linux.SensorLoad15:
			loadAvgSensor.SensorTypeValue = linux.SensorLoad15
		}

		sensors[idx] = loadAvgSensor
	}

	return sensors
}

func parseLoadAvgs(data []byte) ([]string, error) {
	loadAvgsData := bytes.Split(data, []byte(" "))

	if len(loadAvgsData) != 5 { //nolint:mnd
		return nil, ErrParseLoadAvgs
	}

	loadAvgs := make([]string, loadAvgsTotal)

	for idx := range loadAvgs {
		loadAvgs[idx] = string(loadAvgsData[idx][:])
	}

	return loadAvgs, nil
}

func NewLoadAvgWorker(_ context.Context, _ *dbusx.DBusAPI) (*linux.SensorWorker, error) {
	return &linux.SensorWorker{
			Value:    &loadAvgsSensorWorker{loadAvgs: newLoadAvgSensors(), path: filepath.Join(linux.ProcFSRoot, "loadavg")},
			WorkerID: loadAvgsWorkerID,
		},
		nil
}
