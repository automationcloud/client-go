package client

import ()

type JobOutput struct {
	Data interface{} `json:"data"`
}

func (job *Job) GetOutput(key string) (output JobOutput, err error) {
	resp, err := job.apiClient.call("GET", "/jobs/"+job.Id+"/outputs/"+key, nil)
	if err != nil {
		return
	}

	err = ReadBody(resp, &output)
	return
}
