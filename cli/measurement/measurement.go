package measurement

import (
	"fmt"
	"time"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/hostmanager"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
)

// Measurement structure: holds data of a specific measurement
type Measurement struct {
	MeasuredAt time.Time `json:"measuredAt" yaml:"measuredAt"`
	Value      string    `json:"value" yaml:"value"`
	SensorID   string    `json:"sensorId" yaml:"sensorId"`
	DeviceID   string    `json:"deviceId" yaml:"deviceId"`
}

//BeginTempMeasuring function: measures sensor data
func BeginTempMeasuring(hostManager hostmanager.IHostManager, sensorManager sensormanager.ISensorManager, msgChannel chan<- string, doneChannel <-chan bool, unit, format string, delta uint32) {
	var id uint16 = 1

	defer close(msgChannel)

	for {
		select {
		case <-doneChannel:
			return
		default:
			stats, err := hostManager.GetSensorStats()

			if err != nil {
				msgChannel <- constants.ErrMsgRead
				logger.Error.Println("Cannot read sensor data in function <<host.SensorsTemperatures()>>: ", err)
				return
			}

			temp, sensorKey := sensorManager.GetSensorData(&stats, unit, format, id)

			data := CPUTemp{}.NewMeasurement(sensorKey, temp).FormatMeasurement(format)

			msgChannel <- fmt.Sprint(string(*data))

			id++

			time.Sleep(time.Second * time.Duration(delta))
		}
	}
}
