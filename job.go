package client

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

type JobCreationRequest struct {
	ServiceId   string                 `json:"serviceId"`
	Data        map[string]interface{} `json:"input"`
	CallbackUrl string                 `json:"callbackUrl,omitempty"`
}

func (apiClient *ApiClient) CreateJob(jcr JobCreationRequest) (job Job, err error) {
	resp, err := apiClient.call("POST", "/jobs", jcr)

	if err != nil {
		return
	}

	err = readBody(resp, &job)
	if err != nil {
		return
	}

	job.apiClient = apiClient

	return
}

func (job *Job) Cancel() (err error) {
	_, err = job.apiClient.call("POST", "/jobs/"+job.Id+"/cancel", nil)
	return
}

func (job *Job) Fetch() bool {
	resp, err := job.apiClient.call("GET", "/jobs/"+job.Id, nil)
	if err != nil {
		return false
	}

	err = readBody(resp, job)
	if err != nil {
		return false
	}

	return true

}

func (apiClient *ApiClient) FetchJob(id string) (job Job, err error) {
	resp, err := apiClient.call("GET", "/jobs/"+id, nil)
	if err != nil {
		return
	}

	err = readBody(resp, &job)
	if err != nil {
		return
	}

	job.apiClient = apiClient
	return

}
