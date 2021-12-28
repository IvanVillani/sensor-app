package main

import (
	_ "github.com/lib/pq"
	"github.com/seeis/sensor-app/backend/logger"
	"github.com/seeis/sensor-app/backend/manager/osmanager"
	"github.com/seeis/sensor-app/backend/server"
)

func main() {
	logger.SetupLogger(osmanager.OSManager{})

	server.HandleRequests()
}
