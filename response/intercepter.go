package response

import (
	"io"
	"net/http"
)

// Intercepter will intercept http.responseWriter
type Intercepter struct {
	header     http.Header
	StatusCode int

	pr *io.PipeReader
	pw *io.PipeWriter
}

func NewInterceptor() *Intercepter {
	pr, pw := io.Pipe()
	return &Intercepter{pr: pr, pw: pw}
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
	return x.pw.Write(buf)
}

// WriteClose will close writer
func (x Intercepter) WriteClose() error {
	return x.pw.Close()
}

// Read from intercept bytes
func (x Intercepter) Read(buf []byte) (int, error) {
	return x.pr.Read(buf)
}

// ReadClose will close reader
func (x Intercepter) ReadClose() error {
	return x.pr.Close()
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
	_ io.Reader           = &Intercepter{}
	_ http.CloseNotifier  = &Intercepter{}
)
