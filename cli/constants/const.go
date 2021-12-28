package constants

import (
	"io/fs"
	"os"
)

// Constants to be exported to other packages
const (
	// Unit and format
	Farenheit  = "F"
	FormatJSON = "JSON"

	// Sensor
	SensorCPUTemp = "TC0P"

	// Error messages
	ErrMsgParse        = "\nFailed parsing flags...\n"
	ErrMsgRead         = "\nFailed measuring sensor values...\n"
	ErrMsgFlags        = "\n**All of the options must be specified**\n"
	ErrMsgPOST         = "\nFailed executing HTTP.POST request...\n"
	ErrMsgReadRespBody = "\nFailed reading response body from POST request...\n"

	// Info meassages
	InfoMsgEnd = "\nEnd of measurement...\n"

	// Logger
	LogFileName   = "cli.log"
	AppendFlag    = os.O_APPEND
	CreateFlag    = os.O_CREATE
	WriteOnlyFlag = os.O_WRONLY
	FileMod       = fs.FileMode(0666)
)
