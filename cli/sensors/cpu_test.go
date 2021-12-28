package sensors_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seeis/sensor-app/cli/manager/sensormanager"
	"github.com/shirou/gopsutil/host"
)

var _ = Describe("Cpu", func() {
	var (
		statsMatch   *[]host.TemperatureStat
		statsNoMatch *[]host.TemperatureStat
		statsEmpty   *[]host.TemperatureStat

		unitCelsius   = "C"
		unitFarenheit = "F"

		formatJSON = "JSON"
		formatYAML = "YAML"

		id uint16 = 1
	)

	BeforeEach(func() {
		statsMatch = &[]host.TemperatureStat{
			{
				SensorKey:   "TC0P",
				Temperature: 34.00,
			},
			{
				SensorKey:   "TC0F",
				Temperature: 13.00,
			},
		}

		statsNoMatch = &[]host.TemperatureStat{
			{
				SensorKey:   "TC0E",
				Temperature: 34.00,
			},
			{
				SensorKey:   "TC0F",
				Temperature: 13.00,
			},
		}

		statsEmpty = &[]host.TemperatureStat{}
	})

	Describe("Extract sensor data (temperature and sensor key)", func() {
		Context("When match is found / Celsius / JSON", func() {
			It("should match", func() {
				temp, key := sensormanager.SensorManager{}.GetSensorData(statsMatch, unitCelsius, formatJSON, id)

				Expect(temp).To(Equal("34.0°C"))
				Expect(key).To(Equal("TC0P"))
			})
		})

		Context("When match is found / Farenheit / YAML", func() {
			It("should match", func() {
				temp, key := sensormanager.SensorManager{}.GetSensorData(statsMatch, unitFarenheit, formatYAML, id)

				Expect(temp).To(Equal("93.2°F"))
				Expect(key).To(Equal("TC0P"))
			})
		})

		Context("When no match is found", func() {
			It("should match", func() {
				temp, key := sensormanager.SensorManager{}.GetSensorData(statsNoMatch, unitCelsius, formatJSON, id)

				Expect(temp).To(Equal("-"))
				Expect(key).To(Equal("-"))

				temp, key = sensormanager.SensorManager{}.GetSensorData(statsEmpty, unitCelsius, formatJSON, id)

				Expect(temp).To(Equal("-"))
				Expect(key).To(Equal("-"))
			})
		})
	})
})
