package client

type Job struct {
	Id               string `json:"id"`
	ServiceName      string `json:"serviceName"`
	Category         string `json:"category"`
	State            string `json:"state"`
	SessionId        string `json:"sessionId"`
	AwaitingInputKey string `json:"awaitingInputKey,omitempty"`
	apiClient        *ApiClient
}

type JobCreationRequest struct {
	ServiceId   string                 `json:"serviceId"`
	Data        map[string]interface{} `json:"input"`
	CallbackUrl string                 `json:"callbackUrl,omitempty"`
}
