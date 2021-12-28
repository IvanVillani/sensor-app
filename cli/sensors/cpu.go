package sensors

import (
	"fmt"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/shirou/gopsutil/host"
)

// MeasureTemp function: measures cpu temperature
func MeasureTemp(stats *[]host.TemperatureStat, unit, format string, id uint16) (string, string) {
	for _, stat := range *stats {
		if stat.SensorKey != constants.SensorCPUTemp {
			continue
		}

		temperature := convertUnit(unit, stat.Temperature)

		return temperature, stat.SensorKey
	}
	return "-", "-"
}

func convertUnit(unit string, temperature float64) string {
	converted := fmt.Sprintf("%.1f", temperature)

	if unit == constants.Farenheit {
		converted = fmt.Sprintf("%.1f", (temperature*9/5)+32)
	}

	return converted
}
