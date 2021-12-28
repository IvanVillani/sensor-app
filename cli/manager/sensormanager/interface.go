package sensormanager

import "github.com/shirou/gopsutil/host"

// ISensorManager interface: defines method MeasureTemp()
type ISensorManager interface {
	GetSensorData(stats *[]host.TemperatureStat, unit, format string, id uint16) (string, string)
}
