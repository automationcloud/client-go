package client

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type Body struct {
	string
}

func (b Body) Read(p []byte) (n int, err error) {
	if b.string == "fail" {
		return 0, errors.New("Oops")
	}
	bodyBytes := []byte(b.string)
	for i := range bodyBytes {
		p[i] = bodyBytes[i]
	}
	return len(bodyBytes), io.EOF
}

func (b Body) Close() error {
	return nil
}

func TestReadBody(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		b := Body{"{}"}
		res := &http.Response{Body: b}
		var smth interface{}
		err := readBody(res, &smth)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		b := Body{""}
		res := &http.Response{Body: b}
		var smth interface{}
		err := readBody(res, &smth)
		if err.Error() != "unexpected end of JSON input" {
			t.Error("Expected json parsing error")
		}
	})

	t.Run("very unhappy case", func(t *testing.T) {
		b := Body{"fail"}
		res := &http.Response{Body: b}
		var smth interface{}
		err := readBody(res, &smth)
		if err.Error() != "Oops" {
			t.Error("Expected oops error")
		}
	})
}

func TestHttpCall(t *testing.T) {
	t.Run("invalid protocol scheme", func(t *testing.T) {
		apiClient := NewApiClient(&http.Client{}, "").WithBaseURL("")
		_, err := apiClient.call("BADMETHOD", "badurl", nil)
		if !strings.Contains(err.Error(), "unsupported protocol scheme") {
			t.Errorf("Expected 'unsupported protocol scheme' error, got %v", err)
		}
	})

	t.Run("invalid body type", func(t *testing.T) {
		apiClient := NewApiClient(&http.Client{}, "")
		_, err := apiClient.call("GET", "/", make(chan int, 1))
		if !strings.Contains(err.Error(), "unsupported type: chan int") {
			t.Errorf("Expected 'unsupported type: chan int', got %v", err)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		apiClient := NewApiClient(&http.Client{}, "").WithBaseURL("")
		_, err := apiClient.call("/", "https:///", nil)
		if !strings.Contains(err.Error(), "net/http: invalid method \"/\"") {
			t.Errorf("Expected 'invalid method', got %v", err)
		}
	})
}
