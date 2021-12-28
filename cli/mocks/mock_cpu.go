package mocks

import (
	"github.com/shirou/gopsutil/host"
)

// MockCPUSensor struct: implements SesnorManager interface
// Mocked for testing purposes
type MockCPUSensor struct {
	Temperature string
	SensorKey   string
}

// MeasureTemp function: measures cpu temperature
// Mocked for testing purposes
func (mockCPUSensor MockCPUSensor) MeasureTemp(stats *[]host.TemperatureStat, unit, format string, id uint16) (string, string) {
	return mockCPUSensor.Temperature, mockCPUSensor.SensorKey
}
