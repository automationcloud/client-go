package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ApiClient represents api client, use NewApiClient() to create a new client.
type ApiClient struct {
	Client           *http.Client
	SecretKey        string
	baseURL          string
	protocolURL      string
	protocol         *Protocol
	protocolLoadedAt time.Time
}

// NewApiClient creates new ApiClient.
func NewApiClient(httpClient *http.Client, secretKey string) *ApiClient {
	return &ApiClient{
		Client:      httpClient,
		SecretKey:   secretKey,
		baseURL:     "https://api.automationcloud.net",
		protocolURL: "https://protocol.automationcloud.net",
	}
}

// WithProtocolURL allows to alternate location of a protocol.
func (apiClient *ApiClient) WithProtocolURL(url string) *ApiClient {
	apiClient.protocolURL = url
	return apiClient
}

// WithBaseURL allows to alternate location of a api.
func (apiClient *ApiClient) WithBaseURL(url string) *ApiClient {
	apiClient.baseURL = url
	return apiClient
}

func (apiClient *ApiClient) requestProtocol(path string, result interface{}) error {
	_, err := request(
		apiClient.Client,
		"GET",
		apiClient.protocolURL+path,
		"",
		nil,
		result,
	)
	return err
}

func (apiClient *ApiClient) call(method string, path string, payload interface{}, result interface{}) (res *http.Response, err error) {
	return request(
		apiClient.Client,
		method,
		apiClient.baseURL+path,
		apiClient.SecretKey,
		payload,
		result,
	)
}

func request(
	client *http.Client,
	method string,
	url string,
	secretKey string,
	payload interface{},
	result interface{},
) (res *http.Response, err error) {
	var data io.Reader
	if payload != nil {
		json, err := json.Marshal(payload)
		if err != nil {
			return res, err
		}
		data = bytes.NewBuffer(json)
	}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return res, err
	}
	req.Header.Set("Content-Type", "application/json")
	if secretKey != "" {
		req.SetBasicAuth(secretKey, "")
	}
	res, err = client.Do(req)
	if err != nil {
		return res, err
	}

	if 500 <= res.StatusCode && res.StatusCode <= 599 {
		return res, ErrServer
	}

	if res.StatusCode == 400 {
		err = validationError(res)
		return res, err
	}

	if 400 < res.StatusCode && res.StatusCode <= 499 {
		return res, ErrClient
	}

	if result != nil {
		err = readBody(res, &result)
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
