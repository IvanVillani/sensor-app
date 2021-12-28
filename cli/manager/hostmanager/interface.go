package hostmanager

import "github.com/shirou/gopsutil/host"

// IHostManager interface: defines method SensorsTemperatures()
type IHostManager interface {
	GetSensorStats() ([]host.TemperatureStat, error)
}
