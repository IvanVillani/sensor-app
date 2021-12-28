package httpmanager

import (
	"io"
	"net/http"
)

// HTTPManager struct: implements IHTTPManager interface
type HTTPManager struct{}

// Post function: sends post request
func (httpManager HTTPManager) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return http.Post(url, contentType, body)
}
