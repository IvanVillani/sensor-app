package osmanager

import "os"

// IOSManager interface: defines method OpenLogFile()
type IOSManager interface {
	OpenLogFile() (*os.File, error)
}
