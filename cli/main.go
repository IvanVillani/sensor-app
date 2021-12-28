package main

import (
	"github.com/seeis/sensor-app/cli/engine"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
	"github.com/seeis/sensor-app/cli/manager/measurementmanager"
	"github.com/seeis/sensor-app/cli/manager/osmanager"
	"github.com/seeis/sensor-app/cli/manager/writermanager"
)

func main() {
	engine.Start(writermanager.WriterManager{}, osmanager.OSManager{}, flagsmanager.FlagsManager{}, measurementmanager.MeasureManager{})
}
