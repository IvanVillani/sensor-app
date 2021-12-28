package db

import (
	"context"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/lib/pq"
	"github.com/seeis/sensor-app/backend/constants"
	"github.com/seeis/sensor-app/backend/logger"
)

var InfluxClient influxdb2.Client

func InitializeInfluxDBConnection() {
	InfluxAddress := constants.InfluxDBProtocol + "://" +
		os.Getenv("INFLUX_ADDRESS") + ":" +
		constants.InfluxDBPort

	InfluxClient = influxdb2.NewClient(InfluxAddress, constants.AuthToken)

	isReady, err := InfluxClient.Ready(context.Background())

	if err != nil {
		logger.Error.Printf("Connection to InfluxDB failed: %s\n", err)
	}

	if isReady {
		logger.Info.Println(constants.ConnectedToInfluxDBMsg)
	}
}
