package measurement

// IMeasurement interface: creates Measurement types
type IMeasurement interface {
	NewMeasurement() *Measurement
}
