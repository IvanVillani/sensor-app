package measurementmanager

import (
	"github.com/seeis/sensor-app/cli/manager/hostmanager"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
	"github.com/seeis/sensor-app/cli/measurement"
)

// MeasureManager struct: implements IMeasureManager interface
type MeasureManager struct{}

// StartMeasurement function: starts the measurement process
func (measureManager MeasureManager) StartMeasurement(hostManager hostmanager.HostManager, sensorManager sensormanager.SensorManager, msgChannel chan<- string, doneChannel <-chan bool, unit string, format string, delta uint32) {
	measurement.BeginTempMeasuring(hostManager, sensorManager, msgChannel, doneChannel, unit, format, delta)
}
