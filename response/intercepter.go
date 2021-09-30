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

	listener chan []byte
}

func NewInterceptor(ch chan []byte) *Intercepter {
	pr, pw := io.Pipe()
	if ch == nil { // run dummy
		ch = make(chan []byte)
		go func() {
			for b := range ch {
				_ = b
			}
		}()
	}
	return &Intercepter{pr: pr, pw: pw, listener: ch}
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
	x.listener <- buf
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

var (
	// check interface
	_ http.ResponseWriter = &Intercepter{}
	_ io.Writer           = &Intercepter{}
	_ io.Reader           = &Intercepter{}
)
