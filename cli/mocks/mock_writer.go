package mocks

import (
	"bytes"
	"fmt"
	"io"
)

// MockWriter struct: implements SesnorManager interface
// Mocked for testing purposes
type MockWriter struct {
	Output io.Writer
}

// WriteMsg function: writes the data to selected output
// Mocked for testing purposes
func (mockWriter MockWriter) WriteMsg(data string) {
	var b bytes.Buffer

	fmt.Fprint(&b, data)

	b.WriteTo(mockWriter.Output)
}
