package ioutilmanager

import (
	"io"
	"io/ioutil"
)

// IOUtilManager struct: implements IIOUtilManager interface
type IOUtilManager struct{}

// ReadAll function: reads data from response body
func (ioUtilManager IOUtilManager) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
