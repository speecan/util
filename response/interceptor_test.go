package response

import (
	"net/http"
	"testing"

	echo "github.com/labstack/echo/v4"
)

func TestInterceptor(t *testing.T) {
	q := make(chan []byte)
	x := NewInterceptor(q)
	w8 := make(chan struct{})

	t.Run("write response", func(tt *testing.T) {
		tt.Parallel()
		r := echo.NewResponse(x, nil)
		r.WriteHeader(http.StatusOK)
		r.Header().Add("X-Custom-Header-1", "one")
		r.Header().Add("X-Custom-Header-2", "two")
		if _, err := r.Write([]byte("foobar")); err != nil {
			tt.Fatal(err)
		}
		<-w8
	})

	t.Run("assert response", func(tt *testing.T) {
		tt.Parallel()
		b := <-q
		if x.StatusCode != 200 {
			tt.Fatal("status code was unexpected,", x.StatusCode)
		}
		if string(b) != "foobar" {
			tt.Fatal("response body was unexpected,", string(b))
		}
		if x.Header().Get("X-Custom-Header-2") != "two" {
			tt.Fatal("response header was unexpected,", x.Header())
		}
		w8 <- struct{}{}
	})
}
