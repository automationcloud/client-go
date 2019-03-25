package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetProtocol(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"domains":{"A":{}}}`)
	}))
	defer ts.Close()

	now := time.Now()

	ac := NewApiClient(&http.Client{}, "").WithProtocolURL(ts.URL)
	protocol, err := ac.GetProtocol()
	if err != nil {
		panic(err)
	}

	if ac.protocolLoadedAt.Before(now) {
		t.Errorf("Expected protocolLoadedAt to be after %v, but it is: %v", now, ac.protocolLoadedAt)
	}

	_, hasA := protocol.Domains["A"]
	if !hasA {
		t.Errorf("Unexpected protocol")
	}

	if protocol != ac.protocol {
		t.Errorf("Unexpected protocol")
	}

	now = time.Now()
	protocol, _ = ac.GetProtocol()

	if ac.protocolLoadedAt.After(now) {
		t.Errorf("Expected protocolLoadedAt to be after %v, but it is: %v", now, ac.protocolLoadedAt)
	}

	if protocol != ac.protocol {
		t.Errorf("Unexpected protocol")
	}
}
