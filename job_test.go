package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateJob(t *testing.T) {
	t.Run("Invalid job input", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(400)
			w.Header().Add("Content-Type", "application/json")
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

		apiClient := NewApiClient(&http.Client{}, "").WithBaseURL(ts.URL)
		_, err := apiClient.CreateJob(JobCreationRequest{})
		if err == nil {
			t.Error("Expected job creation to fail with 400")
			t.FailNow()
		}

		if !strings.Contains(err.Error(), "validation failed") {
			t.Errorf("Expected validation error, got %v", err)
		}

		if !strings.Contains(err.Error(), "/input/url should be string") {
			t.Errorf("Expected validation error with details, got %v", err)
		}
	})

	t.Run("Valid job input", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(201)
			fmt.Fprintln(w, `{
				"object": "job",
				"id": ":id",
				"createdAt": 1553527953638
			}`)
		}))
		defer ts.Close()

		apiClient := NewApiClient(&http.Client{}, "").WithBaseURL(ts.URL)

		job, err := apiClient.CreateJob(JobCreationRequest{})
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if job.Id != ":id" {
			t.Errorf("Expected job with id :id, got: %v", job.Id)
		}

		if job.apiClient != apiClient {
			t.Error("job should have pointer to apiClient")
		}

		if job.CreatedAt.UTC().Format(time.RFC3339Nano) != "2019-03-25T15:32:33.638Z" {
			t.Errorf("Unexpected job.createdAt: %v", job.CreatedAt.UTC().Format(time.RFC3339Nano))
		}
	})

}

func TestFetch(t *testing.T) {
	status := 200
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, `{
			"object": "job"
		}`)
	}))
	defer ts.Close()
	apiClient := NewApiClient(&http.Client{}, "").WithBaseURL(ts.URL)
	job := &Job{apiClient: apiClient}

	t.Run("happy case", func(t *testing.T) {

		status = 200
		if !job.Fetch() {
			t.Error("job.Fetch() should return true in case of successful request")
		}
	})

	t.Run("unhappy case", func(t *testing.T) {

		status = 500
		if job.Fetch() {
			t.Error("job.Fetch() should return false in case of unsuccessful request")
		}
	})
}

func TestFetchJob(t *testing.T) {
	status := 200
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, `{
			"object": "job"
		}`)
	}))
	defer ts.Close()
	apiClient := NewApiClient(&http.Client{}, "apikey").WithBaseURL(ts.URL)

	t.Run("happy case", func(t *testing.T) {
		status = 200
		job, err := apiClient.FetchJob("id")
		if err != nil {
			t.Error(err)
		}
		if job.apiClient != apiClient {
			t.Error("Expected job to have apiClient link")
		}
	})

	t.Run("unhappy case", func(t *testing.T) {
		status = 404
		_, err := apiClient.FetchJob("id")
		if err == nil {
			t.Error("expected error")
		}
	})
}

func TestCancel(t *testing.T) {
	status := 200
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, `{
			"object": "job"
		}`)
	}))
	defer ts.Close()
	apiClient := NewApiClient(&http.Client{}, "apikey").WithBaseURL(ts.URL)

	t.Run("happy case", func(t *testing.T) {
		job := &Job{apiClient: apiClient, State: "awaitingInput"}
		err := job.Cancel()
		if err != nil {
			t.Errorf("Expect job.Cancel() work without error, got %v", err)
		}
	})

	t.Run("cancel job in final state", func(t *testing.T) {
		job := &Job{apiClient: apiClient, State: "fail"}
		err := job.Cancel()
		if err == nil {
			t.Error("Expect job.Cancel() yield an error, got nothing")
		}
	})
}
