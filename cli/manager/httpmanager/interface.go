package httpmanager

import (
	"io"
	"net/http"
)

// IHTTPManager interface: defines method Post()
type IHTTPManager interface {
	Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
}
