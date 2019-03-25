package client

type JobOutput struct {
	Data      interface{} `json:"data"`
	CreatedAt jsTime      `json:"createdAt"`
	UpdatedAt jsTime      `json:"updatedAt"`
}

// GetOutput loads job output for given key.
func (job *Job) GetOutput(key string) (output JobOutput, err error) {
	resp, err := job.apiClient.call("GET", "/jobs/"+job.Id+"/outputs/"+key, nil)
	if err != nil {
		return
	}

	err = readBody(resp, &output)
	return
}
