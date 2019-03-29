package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOutput(t *testing.T) {
	status := 200
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, `{
			"object": "job-output",
			"data": 13,
			"key": "price"
		}`)
	}))
	defer ts.Close()
	apiClient := NewApiClient(&http.Client{}, "").WithBaseURL(ts.URL)

	job := &Job{apiClient: apiClient}
	output, err := job.GetOutput("price")
	if err != nil {
		t.Errorf("Expect job.Cancel() work without error, got %v", err)
	}

	if output.Data.(float64) != 13 {
		t.Errorf("Expected output data to be 13, got %v", output.Data)
	}
}
