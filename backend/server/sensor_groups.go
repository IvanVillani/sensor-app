package server

// Group specifies different sensor groups
type Group string

// Constants of different sensor groups
const (
	CPUTemp         Group = "CPU_TEMP"
	CPUUsage        Group = "CPU_USAGE"
	MemoryUsage     Group = "MEMORY_USAGE"
	StandartSensors Group = "STANDART_SENSOR"
)
