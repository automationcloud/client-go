package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPreviousJobOutputs(t *testing.T) {
	status := 200
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedPath := "/services/:id/previous-job-outputs"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path to be %v, got: %v", expectedPath, r.URL.Path)
		}
		w.WriteHeader(status)
		fmt.Fprintln(w, `{
			"object": "list",
			"data": [
				{
					"key": "luggageRules",
					"stage": "",
					"data": {
						"title": "Luggage rules",
						"url": "https://example.com/luggage-rules.html"
					},
					"id": "f0d9c116-968c-4dbf-a017-39b92ea729f0",
					"jobId": "701f688d-6689-4f37-8f6d-96675b5d74b3",
					"updatedAt": 1552658147114,
					"createdAt": 1552658147114,
					"variability": 0,
					"object": "previous-job-output"
				}
			]
		}`)
	}))
	defer ts.Close()
	apiClient := NewApiClient(&http.Client{}, "").WithBaseURL(ts.URL)

	t.Run("happy case", func(t *testing.T) {
		status = 200
		_, err := apiClient.ListPreviousJobOutputs(":id")
		if err != nil {
			t.Errorf("Expect apiClient.ListPreviousJobOutputs() work without error, got %v", err)
		}
	})

	t.Run("cancel job in final state", func(t *testing.T) {
		status = 500
		_, err := apiClient.ListPreviousJobOutputs(":id")
		if err == nil {
			t.Error("Expect apiClient.ListPreviousJobOutputs() yield an error, got nothing")
		}
	})
}
