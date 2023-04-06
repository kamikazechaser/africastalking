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

// New returns an instance of an Africa's Talking client reusbale across different products.
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

// SetHTTPClient can be used to override the default client with a custom set one.
func (at *atClient) SetHTTPClient(httpClient *http.Client) *atClient {
	at.httpClient = httpClient

	return at
}

// setDefaultHeaders sets the standard headers required by the Africa's Talking API.
func (at *atClient) setDefaultHeaders(req *http.Request) *http.Request {
	req.Header.Set("apiKey", at.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
