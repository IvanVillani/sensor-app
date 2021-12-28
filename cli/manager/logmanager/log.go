package logmanager

import (
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/osmanager"
)

// LogManager struct: implements ILogManager interface
type LogManager struct{}

// SetupLog function: sets up the logger
func (logManager LogManager) SetupLog() {
	logger.SetupLogger(osmanager.OSManager{})
}
