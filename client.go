package africastalking

import (
	"net/http"
	"time"
)

const (
	baseLiveEndpoint    = "https://api.africastalking.com/version1/"
	baseSandboxEndpoint = "https://api.sandbox.africastalking.com/version1/"
)

type (
	atClient struct {
		apiKey     string
		endpoint   string
		httpClient *http.Client
		username   string
	}
)

func New(apiKey string, username string, sandbox bool) *atClient {
	atClient := &atClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		username: username,
	}

	if sandbox {
		atClient.endpoint = baseSandboxEndpoint
	} else {
		atClient.endpoint = baseLiveEndpoint
	}

	return atClient
}

func (at *atClient) SetHTTPClient(httpClient *http.Client) *atClient {
	at.httpClient = httpClient

	return at
}

func (at *atClient) setDefaultHeaders(req *http.Request) *http.Request {
	req.Header.Set("apiKey", at.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
