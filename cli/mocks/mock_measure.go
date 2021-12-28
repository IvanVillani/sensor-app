package mocks

import (
	"errors"
	"time"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/hostmanager"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
)

// MockMeasure struct: implements MeasureManager interface
// Mocked for testing purposes
type MockMeasure struct {
	Uninterrupted bool
}

// StartMeasurement function: starts the measurement process
// Mocked for testing purposes
func (mockMeasure MockMeasure) StartMeasurement(hostManager hostmanager.HostManager, sensorManager sensormanager.SensorManager, msgChannel chan<- string, doneChannel <-chan bool, unit string, format string, delta uint32) {
	var id uint16 = 1

	defer close(msgChannel)

	if mockMeasure.Uninterrupted {
		for {
			select {
			case <-doneChannel:
				return
			default:
				msgChannel <- "Mocked measurement"

				id++

				time.Sleep(time.Second * time.Duration(delta))
			}
		}
	} else {
		for {
			select {
			case <-doneChannel:
				return
			default:
				msgChannel <- constants.ErrMsgRead
				logger.Error.Println("Cannot read sensor data in function <<host.SensorsTemperatures()>>: ", errors.New("mocked error").Error())
				return
			}
		}
	}
}
