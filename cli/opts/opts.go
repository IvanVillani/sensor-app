package opts

import (
	"github.com/jessevdk/go-flags"
	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
)

// ParseOpts function: parses CLI arguments into options
func ParseOpts(flagsManager flagsmanager.IFlagsManager) (string, error) {
	if _, err := flagsManager.ParseFlags(); err != nil {
		switch flagsErr := err.(type) {
		case *flags.Error:
			if flagsErr.Type == flags.ErrHelp {
				logger.Info.Println("Used \"--help\" option.")
				return constants.ErrMsgFlags, err
			}
			logger.Warning.Println("Required flags were not specified: ", err)
			return "", err
		default:
			logger.Error.Println("Cannot parse arguments in function <<flags.Parse()>>: ", err)
			return constants.ErrMsgParse, err
		}
	}
	return "", nil
}
