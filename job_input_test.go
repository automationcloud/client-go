package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateInput(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(400)
			fmt.Fprintln(w, `{ "object": "error",
				"code": "badRequest",
				"name": "ValidationError",
				"message": "Job validation failed",
				"details": {
					"name": "Job",
					"messages": [
						"/input/url should be string"
					]
				},
				"stack": "Error: Job validation failed\n    at Object.create (/src/app/services/job/create.js:142:15)\n    at process._tickCallback (internal/process/next_tick.js:68:7)"
			}`)
		}))
		defer ts.Close()

		client := &http.Client{}
		// witness.DebugClient(client)
		apiClient := NewApiClient(client, "").WithBaseURL(ts.URL)
		job := Job{Id: ":id"}
		job.apiClient = apiClient
		_, err := job.CreateInput(InputCreationRequest{})
		if err == nil {
			t.Error("Expected input creation to fail with 400")
			t.FailNow()
		}

		if !strings.Contains(err.Error(), "validation failed") {
			t.Errorf("Expected validation error, got %v", err)
		}
	})

}
