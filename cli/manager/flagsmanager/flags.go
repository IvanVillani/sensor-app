package flagsmanager

import (
	"github.com/jessevdk/go-flags"
)

// Opts structure: holds all flag input options
var Opts struct {
	// Unit type of the measurement (Celsius or Farenheit)
	Unit string `long:"unit" required:"true" choice:"C" choice:"F" description:"Unit type of the measurement (Celsius or Farenheit)"`

	// Delta duration in seconds between two successive sensor measurements
	DeltaDuration uint32 `long:"delta_duration" required:"true" description:"Delta duration in seconds between two successive sensor measurements"`

	// Total duration in seconds for all sensor measurements
	TotalDuration uint64 `long:"total_duration" required:"true" description:"Total duration in seconds for all sensor measurements"`

	// Format the output temperature readings (JSON or YAML)
	Format string `long:"format" choice:"JSON" required:"true" choice:"YAML" description:"Format the output temperature readings (JSON or YAML)"`

	// Server URL for saving measurements
	WebHookURL string `long:"web_hook_url" description:"Server URL for saving measurements"`
}

// FlagsManager struct: implements IFlagsManager interface
type FlagsManager struct{}

// ParseFlags function: parse flags from input
func (flagsManager FlagsManager) ParseFlags() ([]string, error) {
	return flags.Parse(&Opts)
}

func WebHookURLProvided() bool {
	return Opts.WebHookURL != ""
}
