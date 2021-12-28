package mocks

import (
	"io"

	"github.com/seeis/sensor-app/cli/logger"
)

// MockLogger struct: implements LogManager interface
// Mocked for testing purposes
type MockLogger struct {
	Writer io.Writer
}

// SetupLog function: sets up the logger
// Mocked for testing purposes
func (mockLogger MockLogger) SetupLog() {
	logger.MockSetupLogger(mockLogger.Writer)
}
