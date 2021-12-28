package mocks

import (
	"github.com/shirou/gopsutil/host"
)

// MockHost struct: implements HostManager interface
// Mocked for testing purposes
type MockHost struct {
	Stats []host.TemperatureStat
	Error error
}

// GetSensorStats function: get sensor stats
// Mocked for testing purposes
func (mockHost MockHost) GetSensorStats() ([]host.TemperatureStat, error) {
	return mockHost.Stats, mockHost.Error
}
