package client

import (
	"fmt"
)

type InputCreationRequest struct {
	Key   string      `json:"key"`
	Stage string      `json:"stage"`
	Data  interface{} `json:"data"`
}

type JobInput struct {
	Key   string      `json:"key"`
	Stage string      `json:"stage"`
	Data  interface{} `json:"data"`
}

func (job *Job) CreateInput(inputCreationRequest InputCreationRequest) (input JobInput, err error) {
	fmt.Println("Create input", inputCreationRequest)
	resp, err := job.apiClient.call("POST", "/jobs/"+job.Id+"/inputs", inputCreationRequest)
	if err != nil {
		return
	}

	err = ReadBody(resp, &input)
	fmt.Println(err, input)
	return
}
