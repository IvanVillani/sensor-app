package ioutilmanager

import (
	"io"
)

// IIOUtilManager interface: defines method ReadAll()
type IIOUtilManager interface {
	ReadAll(r io.Reader) ([]byte, error)
}
