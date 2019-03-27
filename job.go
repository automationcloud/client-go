package client

import (
	"errors"
	"fmt"
)

// Job represents automation cloud job object.
// It used to control automation flow: provide inputs, consume outputs, watch state.
type Job struct {
	Id               string `json:"id"`
	ServiceName      string `json:"serviceName"`
	Category         string `json:"category"`
	State            string `json:"state"`
	SessionId        string `json:"sessionId"`
	AwaitingInputKey string `json:"awaitingInputKey,omitempty"`
	CreatedAt        jsTime `json:"createdAt"`
	UpdatedAt        jsTime `json:"updatedAt"`
	apiClient        *ApiClient
}

// JobCreationRequest describes a request to create a job.
type JobCreationRequest struct {
	ServiceId   string                 `json:"serviceId"`
	Data        map[string]interface{} `json:"input"`
	CallbackUrl string                 `json:"callbackUrl,omitempty"`
}

// CreateJob makes an http request to run a job.
// It may return ValidationError in case if invalid data provided.
func (apiClient *ApiClient) CreateJob(jcr JobCreationRequest) (job Job, err error) {
	_, err = apiClient.call("POST", "/jobs", jcr, &job)

	if err != nil {
		return
	}

	job.apiClient = apiClient

	return
}

// Cancel sends http requst to cancel job
// Yields an error in case of attempt to cancel job not in "awaitingInput" state
func (job *Job) Cancel() (err error) {
	if job.State != "awaitingInput" {
		return errors.New(fmt.Sprintf("can not cancel job in %v state", job.State))
	}
	_, err = job.apiClient.call("POST", "/jobs/"+job.Id+"/cancel", nil, nil)
	return
}

// Fetch loads updated version of a job, returns true if request was successful.
func (job *Job) Fetch() bool {
	_, err := job.apiClient.call("GET", "/jobs/"+job.Id, nil, job)
	if err != nil {
		return false
	}

	return true

}

// FetchJob requests job and returns it as a result.
func (apiClient *ApiClient) FetchJob(id string) (job Job, err error) {
	_, err = apiClient.call("GET", "/jobs/"+id, nil, &job)
	job.apiClient = apiClient
	return
}
