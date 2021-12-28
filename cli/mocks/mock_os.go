package mocks

import (
	"errors"
	"os"

	"github.com/seeis/sensor-app/cli/constants"
)

// MockOS struct: implements OSManager interface
// Mocked for testing purposes
type MockOS struct {
	MockedWithErr bool
}

// OpenLogFile function: opens file for logging
// Mocked for testing purposes
func (mockOS MockOS) OpenLogFile() (*os.File, error) {
	if mockOS.MockedWithErr {
		return &os.File{}, errors.New("can't open/create the file")
	}
	return os.OpenFile(constants.LogFileName, constants.AppendFlag|constants.CreateFlag|constants.WriteOnlyFlag, constants.FileMod)
}
