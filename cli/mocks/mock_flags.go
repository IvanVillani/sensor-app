package mocks

import (
	"errors"

	"github.com/jessevdk/go-flags"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
)

// MockFlags struct: implements FlagsManager interface
// Mocked for testing purposes
type MockFlags struct {
	MockedWithErr    bool
	ErrorIsFromFlags bool
	ErrorIsHelpError bool
}

// ParseFlags function: parse flags from input
// Mocked for testing purposes
func (mockFlags MockFlags) ParseFlags() ([]string, error) {
	slice := []string{"", ""}

	if mockFlags.MockedWithErr {
		if mockFlags.ErrorIsFromFlags {
			if mockFlags.ErrorIsHelpError {
				return slice, &flags.Error{
					Type:    flags.ErrHelp,
					Message: "Error help message from MockFlags",
				}
			}
			return slice, &flags.Error{
				Type:    flags.ErrCommandRequired,
				Message: "Error flag message from MockFlags",
			}
		}
		return slice, errors.New("unknown error")
	}
	flagsmanager.Opts.Unit = "C"
	flagsmanager.Opts.Format = "JSON"
	flagsmanager.Opts.DeltaDuration = 2
	flagsmanager.Opts.TotalDuration = 4

	return slice, nil
}
