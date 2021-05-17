package response

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	echo "github.com/labstack/echo/v4"
)

func TestInterceptor(t *testing.T) {
	// q := make(chan []byte)
	x := NewInterceptor()
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
		x.WriteClose()
		<-w8
	})

	t.Run("assert response", func(tt *testing.T) {
		tt.Parallel()
		defer x.ReadClose()
		b, err := ioutil.ReadAll(x)
		if err != nil {
			tt.Fatal(err)
		}
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

func TestInterceptorWithPipe(t *testing.T) {
	// q := make(chan []byte, 10240)
	// respr, respw := io.Pipe()
	x := NewInterceptor()
	w8 := make(chan struct{})
	pr, pw := io.Pipe()

	t.Run("write to pipe", func(tt *testing.T) {
		tt.Parallel()
		defer pw.Close()
		sr := strings.NewReader(strings.Repeat("foobar12", 128*10)) // 8bytes * 128 = 1kb
		_, err := io.Copy(pw, sr)
		if err != nil {
			tt.Fatal(err)
		}
	})

	t.Run("write response", func(tt *testing.T) {
		tt.Parallel()
		r := echo.NewResponse(x, nil)
		r.Header().Add("X-Custom-Header-1", "one")
		r.Header().Add("X-Custom-Header-2", "two")
		r.WriteHeader(http.StatusOK)
		n, err := io.Copy(r, pr)
		if err != nil {
			tt.Fatal(n, err)
		}
		pr.Close()
		x.WriteClose()
		<-w8
	})

	t.Run("assert response", func(tt *testing.T) {
		tt.Parallel()
		// b := <-q
		defer x.ReadClose()
		b, err := ioutil.ReadAll(x)
		if err != nil {
			tt.Fatal(err)
		}
		if x.StatusCode != 200 {
			tt.Fatal("status code was unexpected,", x.StatusCode)
		}
		if string(b) != strings.Repeat("foobar12", 128*10) {
			tt.Fatal("response body was unexpected,", string(b))
		}
		if x.Header().Get("X-Custom-Header-2") != "two" {
			tt.Fatal("response header was unexpected,", x.Header())
		}
		w8 <- struct{}{}
	})
}

func TestHTTPServer(t *testing.T) {
	e := echo.New()
	mid := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			orgWriter := c.Response().Writer
			w := NewInterceptor()
			c.Response().Writer = w
			w8 := make(chan error)
			go func() {
				body, err := ioutil.ReadAll(w)
				if err != nil {
					w8 <- err
					return
				}
				//
				fmt.Println(string(body)) // Save to database
				c.Response().Writer = orgWriter
				for k, v := range w.Header() {
					for _, s := range v {
						c.Response().Header().Add(k, s)
					}
				}
				c.Response().WriteHeader(w.StatusCode)
				c.Response().Write(body)
				w8 <- nil
			}()
			if err := next(c); err != nil {
				return err
			}
			return <-w8
		}
	}
	e.GET("/dummy", func(c echo.Context) error {
		sr := strings.NewReader(strings.Repeat("foobar12", 128*10))
		return c.Stream(http.StatusOK, "application/octet-stream", sr)
	}, mid)
	s := httptest.NewServer(e)
	resp, err := s.Client().Get(s.URL + "/dummy")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(string(b))
}
