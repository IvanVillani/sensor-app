package mocks

type MockReadCloser struct{}

func (mockReadCloser MockReadCloser) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (mockReadCloser MockReadCloser) Close() error {
	return nil
}
