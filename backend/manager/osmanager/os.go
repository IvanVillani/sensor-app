package osmanager

import (
	"os"

	"github.com/seeis/sensor-app/backend/constants"
)

// OSManager struct: implements IOSManager interface
type OSManager struct{}

// OpenLogFile function: opens file for logging
func (osManager OSManager) OpenLogFile() (*os.File, error) {
	return os.OpenFile(constants.LogFileName, constants.AppendFlag|constants.CreateFlag|constants.WriteOnlyFlag, constants.FileMod)
}
