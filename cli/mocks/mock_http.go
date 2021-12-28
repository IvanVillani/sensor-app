package mocks

import (
	"io"
	"net/http"
)

// MockHttp struct: implements IHttpManager interface
// Mocked for testing purposes
type MockHttp struct {
	Status string
	Error  error
}

// Post function: executes POST request
// Mocked for testing purposes
func (mockHTTP MockHttp) Post(url string, contentType string, body io.Reader) (resp *http.Response, err error) {
	return &http.Response{
		Status:     mockHTTP.Status,
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Body:       MockReadCloser{},
		Close:      false,
	}, mockHTTP.Error
}
