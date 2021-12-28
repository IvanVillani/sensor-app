package measurement_test

import (
	"errors"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shirou/gopsutil/host"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
	"github.com/seeis/sensor-app/cli/manager/writermanager"
	"github.com/seeis/sensor-app/cli/measurement"
	"github.com/seeis/sensor-app/cli/mocks"
)

var _ = Describe("Measure", func() {
	var (
		unit   = "C"
		format = "JSON"

		formatJSON = "JSON"
		formatYAML = "YAML"

		msgChannel  chan string
		doneChannel chan bool

		delta uint32 = 2
	)

	Describe("Measurement process is running", func() {
		Context("When match is found", func() {
			It("should post to the channel the right temperature", func() {
				hostManager := mocks.MockHost{
					Stats: []host.TemperatureStat{
						{
							Temperature: 99.9,
							SensorKey:   "TC0P",
						},
					},
					Error: nil,
				}

				sensorManager := sensormanager.SensorManager{}

				msgChannel = make(chan string)
				doneChannel = make(chan bool)
				defer close(doneChannel)
				untilChannel := time.After(time.Duration(1) * time.Second)

				go measurement.BeginTempMeasuring(hostManager, sensorManager, msgChannel, doneChannel, unit, format, delta)

				for {
					select {
					case <-untilChannel:
						doneChannel <- true
						writermanager.WriterManager{}.WriteMsg(constants.InfoMsgEnd)
						return
					case m := <-msgChannel:
						Expect(m).To(ContainSubstring("99.9°C"))
					}
				}
			})
		})

		Context("When no match is found", func() {
			It("should post to the channel empty temperature value", func() {
				hostManager := mocks.MockHost{
					Stats: []host.TemperatureStat{
						{
							Temperature: 99.9,
							SensorKey:   "TC0F",
						},
					},
					Error: nil,
				}

				sensorManager := sensormanager.SensorManager{}

				msgChannel = make(chan string)
				doneChannel = make(chan bool)
				defer close(doneChannel)
				untilChannel := time.After(time.Duration(1) * time.Second)

				go measurement.BeginTempMeasuring(hostManager, sensorManager, msgChannel, doneChannel, unit, format, delta)

				for {
					select {
					case <-untilChannel:
						doneChannel <- true
						writermanager.WriterManager{}.WriteMsg(constants.InfoMsgEnd)
						return
					case m := <-msgChannel:
						Expect(m).To(ContainSubstring("-"))
					}
				}
			})
		})

		Context("When error is returned", func() {
			It("should post to the channel terminal error message", func() {
				mocks.MockLogger{
					Writer: os.Stdout,
				}.SetupLog()

				hostManager := mocks.MockHost{
					Stats: []host.TemperatureStat{
						{
							Temperature: 99.9,
							SensorKey:   "TC0F",
						},
					},
					Error: errors.New("Something went wrong..."),
				}

				sensorManager := sensormanager.SensorManager{}

				msgChannel = make(chan string)
				doneChannel = make(chan bool)
				defer close(doneChannel)
				untilChannel := time.After(time.Duration(1) * time.Second)

				go measurement.BeginTempMeasuring(hostManager, sensorManager, msgChannel, doneChannel, unit, format, delta)

				for {
					select {
					case <-untilChannel:
						doneChannel <- true
						writermanager.WriterManager{}.WriteMsg(constants.InfoMsgEnd)
						return
					case m := <-msgChannel:
						if m == constants.ErrMsgRead {
							Expect(m).To(Equal(constants.ErrMsgRead))
							return
						}
					}
				}
			})
		})

		Context("Check formatting data function", func() {
			It("should format the data in the correct format", func() {
				measurementData := measurement.CPUTemp{}.NewMeasurement("TC0P", "99.9°C").FormatMeasurement(formatJSON)

				result := string(*measurementData)

				Expect(result).To(ContainSubstring("{\"measuredAt\":\""))
				Expect(result).To(ContainSubstring("\",\"value\":\"99.9°C\",\"sensorId\":\"TC0P\",\"deviceId\":\"MBP16-W\"}"))

				measurementData = measurement.CPUTemp{}.NewMeasurement("TC0P", "99.9°C").FormatMeasurement(formatYAML)

				result = string(*measurementData)

				Expect(result).To(ContainSubstring("value: 99.9°C\nsensorId: TC0P\ndeviceId: MBP16-W"))
			})
		})
	})
})
