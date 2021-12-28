package measurementmanager

import (
	"github.com/seeis/sensor-app/cli/manager/hostmanager"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
)

// IMeasureManager interface: defines method StartMeasurement()
type IMeasureManager interface {
	StartMeasurement(hostManager hostmanager.HostManager, sensorManager sensormanager.SensorManager, msgChannel chan<- string, doneChannel <-chan bool, unit string, format string, delta uint32)
}
