package constants

import (
	"io/fs"
	"os"
)

// Constants to be exported to other packages
const (
	// PostgreSQL
	PostgresDriverName = "postgres"

	// Database connection
	InfluxDBProtocol    = "http"
	InfluxDBPort        = "8086"
	PostgresSSLDisabled = "?sslmode=disable"

	// InfluxDB
	AuthToken  = "ZqwqPwOAEyQ-d7KiaJIVmrW6hsXAscJh1jxTQJ7QUuIEwTBNpJoTxCtUP84Z-ErH3qUiQqVOjX9BXtdjvAVedA==13"
	OrgName    = "sensor-app-org"
	BucketName = "sensor-app-bucket"

	// Others
	TypeOfMeasurement = "cpu"

	// Info messages
	ConnectedToPostgreSQLMsg = "Successfully connected to PostgreSQL!"
	ConnectedToInfluxDBMsg   = "Successfully connected to InfluxDB!"

	// Logger
	LogFileName   = "backend.log"
	AppendFlag    = os.O_APPEND
	CreateFlag    = os.O_CREATE
	WriteOnlyFlag = os.O_WRONLY
	FileMod       = fs.FileMode(0666)
)
