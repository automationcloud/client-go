package client

import "fmt"

type PreviousJobOutput struct {
	Key         string      `json:"key"`
	Stage       string      `json:"stage"`
	Data        interface{} `json:"data"`
	Id          string      `json:"id"`
	JobId       string      `json:"jobId"`
	UpdatedAt   jsTime      `json:"updatedAt"`
	CreatedAt   jsTime      `json:"createdAt"`
	Variability float64     `json:"variability"`
}

// ListPreviousJobOutputs fetches previous outputs for a service.
func (apiClient *ApiClient) ListPreviousJobOutputs(serviceId string) ([]PreviousJobOutput, error) {
	var body struct {
		Data []PreviousJobOutput `json:"data"`
	}

	_, err := apiClient.call(
		"GET",
		fmt.Sprintf("/services/%s/previous-job-outputs", serviceId),
		nil,
		&body,
	)
	if err != nil {
		return nil, err
	}

	return body.Data, nil
}
