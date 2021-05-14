package response

import (
	"io"
	"net/http"
)

// Intercepter will intercept http.responseWriter
type Intercepter struct {
	header     http.Header
	Listener   chan []byte
	StatusCode int
}

func NewInterceptor(q chan []byte) *Intercepter {
	return &Intercepter{Listener: q}
}

// Header is returning header map
func (x *Intercepter) Header() http.Header {
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

// WriteHeader will set own status code
func (x *Intercepter) WriteHeader(statusCode int) {
	x.StatusCode = statusCode
}

// CloseNotify is dummy
func (Intercepter) CloseNotify() <-chan bool {
	return make(chan bool)
}

var (
	// check interface
	_ http.ResponseWriter = &Intercepter{}
	_ io.Writer           = &Intercepter{}
	_ http.CloseNotifier  = &Intercepter{}
)
