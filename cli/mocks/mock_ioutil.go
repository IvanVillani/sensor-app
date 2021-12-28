package mocks

import (
	"io"
)

// MockIOUtil struct: implements IIOUtilManager interface
// Mocked for testing purposes
type MockIOUtil struct {
	Msg   string
	Error error
}

// ReadAll function: reads data from response body
// Mocked for testing purposes
func (mockIOUtil MockIOUtil) ReadAll(r io.Reader) ([]byte, error) {
	return []byte(mockIOUtil.Msg), mockIOUtil.Error
}
