package engine

import (
	"time"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
	"github.com/seeis/sensor-app/cli/manager/hostmanager"
	"github.com/seeis/sensor-app/cli/manager/measurementmanager"
	"github.com/seeis/sensor-app/cli/manager/osmanager"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
	"github.com/seeis/sensor-app/cli/manager/writermanager"
	"github.com/seeis/sensor-app/cli/opts"
)

//Start function: prepares and starts the measurement program
func Start(writerManager writermanager.IWriterManager, osManager osmanager.IOSManager, flagsManager flagsmanager.IFlagsManager, measureManager measurementmanager.IMeasureManager) {
	logger.SetupLogger(osManager)

	msg, err := opts.ParseOpts(flagsManager)

	if err == nil {
		msgChannel := make(chan string)
		doneChannel := make(chan bool)
		defer close(doneChannel)
		untilChannel := time.After(time.Duration(flagsmanager.Opts.TotalDuration) * time.Second)

		go measureManager.StartMeasurement(hostmanager.HostManager{}, sensormanager.SensorManager{}, msgChannel, doneChannel, flagsmanager.Opts.Unit, flagsmanager.Opts.Format, flagsmanager.Opts.DeltaDuration)

		for {
			select {
			case <-untilChannel:
				doneChannel <- true
				writerManager.WriteMsg(constants.InfoMsgEnd)
				return
			case m := <-msgChannel:
				if m == constants.ErrMsgRead {
					writerManager.WriteMsg(m)
					return
				}
				writerManager.WriteMsg(m)
			}
		}
	}

	writerManager.WriteMsg(msg)
}
