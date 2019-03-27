package client

// JobOutput represents a job output.
type JobOutput struct {
	Data      interface{} `json:"data"`
	CreatedAt jsTime      `json:"createdAt"`
	UpdatedAt jsTime      `json:"updatedAt"`
}

// GetOutput loads job output for given key.
func (job *Job) GetOutput(key string) (output JobOutput, err error) {
	_, err = job.apiClient.call("GET", "/jobs/"+job.Id+"/outputs/"+key, nil, output)
	return
}
