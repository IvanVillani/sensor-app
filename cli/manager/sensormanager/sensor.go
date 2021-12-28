package sensormanager

import (
	"github.com/seeis/sensor-app/cli/sensors"
	"github.com/shirou/gopsutil/host"
)

// SensorManager struct: implements ISesnorManager interface
type SensorManager struct{}

// GetSensorData function: measures cpu temperature
func (sensorManager SensorManager) GetSensorData(stats *[]host.TemperatureStat, unit, format string, id uint16) (string, string) {
	return sensors.MeasureTemp(stats, unit, format, id)
}
