package hostmanager

import (
	"github.com/shirou/gopsutil/host"
)

// HostManager struct: implements IHostManager interface
type HostManager struct{}

// GetSensorStats function: get sensor stats
func (hostManager HostManager) GetSensorStats() ([]host.TemperatureStat, error) {
	return host.SensorsTemperatures()
}
