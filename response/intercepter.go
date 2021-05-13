package response

import "net/http"

// Intercepter will intercept response.writer
type Intercepter struct {
	header   http.Header
	Listener chan []byte
}

// Header is returning header map
func (x Intercepter) Header() http.Header {
	if x.header == nil {
		x.header = map[string][]string{}
	}
	return x.header
}

// Write will intercept bytes
func (x Intercepter) Write(buf []byte) (int, error) {
	x.Listener <- buf
	return 0, nil
}

// WriteHeader is dummy
func (Intercepter) WriteHeader(x int) {
}

// CloseNotify is dummy
func (Intercepter) CloseNotify() <-chan bool {
	return make(chan bool)
}
