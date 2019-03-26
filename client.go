package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiClient struct {
	Client           *http.Client
	SecretKey        string
	baseURL          string
	protocolURL      string
	protocol         *Protocol
	protocolLoadedAt time.Time
}

func NewApiClient(httpClient *http.Client, secretKey string) *ApiClient {
	return &ApiClient{
		Client:      httpClient,
		SecretKey:   secretKey,
		baseURL:     "https://api.automationcloud.net",
		protocolURL: "https://protocol.automationcloud.net/schema.json",
	}
}

func (apiClient *ApiClient) WithProtocolURL(url string) *ApiClient {
	apiClient.protocolURL = url
	return apiClient
}

func (apiClient *ApiClient) WithBaseURL(url string) *ApiClient {
	apiClient.baseURL = url
	return apiClient
}

func (apiClient *ApiClient) call(method string, path string, payload interface{}) (res *http.Response, err error) {
	json, err := json.Marshal(payload)
	req, err := http.NewRequest(method, apiClient.baseURL+path, bytes.NewBuffer(json))
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiClient.SecretKey, "")
	res, err = apiClient.Client.Do(req)
	if err != nil {
		return res, err
	}

	if 500 <= res.StatusCode && res.StatusCode <= 599 {
		return res, ServerError
	}

	if res.StatusCode == 400 {
		err = validationError(res)
		return res, err
	}

	if 400 < res.StatusCode && res.StatusCode <= 499 {
		return res, ClientError
	}

	// fmt.Println(method, path, res.Status)
	return res, err
}

func readBody(res *http.Response, data interface{}) (err error) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, data)
	if err != nil {
		return
	}
	return
}
