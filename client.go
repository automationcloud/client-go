package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiClient struct {
	Client    *http.Client
	BaseUrl   string
	SecretKey string
}

func (apiClient *ApiClient) call(method string, path string, payload interface{}) (res *http.Response, err error) {
	json, err := json.Marshal(payload)
	req, err := http.NewRequest(method, apiClient.BaseUrl+path, bytes.NewBuffer(json))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiClient.SecretKey, "")
	res, err = apiClient.Client.Do(req)
	if err != nil {
		return
	}
	fmt.Println(method, path, res.Status)
	return
}

func (apiClient *ApiClient) CreateJob(serviceId string, data map[string]interface{}, callbackUrl string) (job Job, err error) {
	resp, err := apiClient.call(
		"POST",
		"/jobs",
		JobCreationRequest{ServiceId: serviceId, Data: data, CallbackUrl: callbackUrl},
	)

	if err != nil {
		return
	}

	err = ReadBody(resp, &job)
	job.apiClient = apiClient
	return
}

func ReadBody(res *http.Response, data interface{}) (err error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	// fmt.Println(string(body))
	err = json.Unmarshal(body, data)
	if err != nil {
		return
	}
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

	err = ReadBody(resp, job)
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

	err = ReadBody(resp, &job)
	if err != nil {
		return
	}

	job.apiClient = apiClient
	return

}

func (apiClient *ApiClient) Get(url string, data interface{}) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	res, err := apiClient.Client.Do(req)
	if err != nil {
		return
	}

	err = ReadBody(res, &data)
	return

}
