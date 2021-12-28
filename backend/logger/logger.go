package logger

import (
	"log"
	"os"

	"github.com/seeis/sensor-app/backend/manager/osmanager"
)

var (
	// Info log: Just informative
	Info *log.Logger

	// Warning log: Be concerned
	Warning *log.Logger

	// Error log: Important case
	Error *log.Logger

	// Fatal log: Critical case
	Fatal *log.Logger
)

// SetupLogger function: initializes logger levels and set output Writer
func SetupLogger(osManager osmanager.IOSManager) {
	file, err := osManager.OpenLogFile()

	wr := file

	if err != nil {
		wr = os.Stderr
	}

	Info = log.New(wr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(wr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(wr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(wr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

}
