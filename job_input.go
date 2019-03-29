package client

// inputCreationRequest describes parameters for request to create input.
type inputCreationRequest struct {
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

// CreateInput creates an input for job.
func (job *Job) CreateInput(data interface{}) (input JobInput, err error) {
	icr := inputCreationRequest{
		Key:   job.AwaitingInputKey,
		Stage: job.AwaitingInputStage,
		Data:  data,
	}
	_, err = job.apiClient.call("POST", "/jobs/"+job.Id+"/inputs", icr, &input)
	return
}
