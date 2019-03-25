package client

import (
	"fmt"
)

// InputCreationRequest
type InputCreationRequest struct {
	Key   string      `json:"key"`
	Stage string      `json:"stage,omitempty"`
	Data  interface{} `json:"data"`
}

// JobInput represents a job input.
type JobInput struct {
	Key       string      `json:"key"`
	Stage     string      `json:"stage"`
	Data      interface{} `json:"data"`
	CreatedAt int         `json:"createdAt"`
}

// CreateInput creates an input for job
func (job *Job) CreateInput(inputCreationRequest InputCreationRequest) (input JobInput, err error) {
	resp, err := job.apiClient.call("POST", "/jobs/"+job.Id+"/inputs", inputCreationRequest)
	if err != nil {
		return
	}

	err = readBody(resp, &input)
	fmt.Println(err, input)
	return
}
