package measurement

import (
	"encoding/json"
	"time"

	"github.com/seeis/sensor-app/cli/constants"
	"gopkg.in/yaml.v2"
)

// CPUTemp structure: holds data of a specific measurement
type CPUTemp struct{}

// NewMeasurement function: creates Measurement struct
func (cpuTemp CPUTemp) NewMeasurement(sensorKey, temp string) *Measurement {
	msmnt := Measurement{
		MeasuredAt: time.Now(),
		Value:      temp,
		SensorID:   sensorKey,
		DeviceID:   "MBP13",
	}

	return &msmnt
}

// FormatMeasurement function: formats data into JSON or YAML
func (msmnt *Measurement) FormatMeasurement(format string) *[]byte {
	var data []byte

	if format == constants.FormatJSON {
		j, _ := json.Marshal(*msmnt)
		data = j
		return &data
	}

	y, _ := yaml.Marshal(msmnt)
	data = y

	return &data
}
