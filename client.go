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
	AtClient struct {
		apiKey     string
		endpoint   string
		httpClient *http.Client
		username   string
	}
)

// New returns an instance of an Africa's Talking client reusbale across different products.
func New(apiKey string, username string, sandbox bool) *AtClient {
	AtClient := &AtClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		username: username,
	}

	if sandbox {
		AtClient.endpoint = baseSandboxEndpoint
	} else {
		AtClient.endpoint = baseLiveEndpoint
	}

	return AtClient
}

// SetHTTPClient can be used to override the default client with a custom set one.
func (at *AtClient) SetHTTPClient(httpClient *http.Client) *AtClient {
	at.httpClient = httpClient

	return at
}

// setDefaultHeaders sets the standard headers required by the Africa's Talking API.
func (at *AtClient) setDefaultHeaders(req *http.Request) *http.Request {
	req.Header.Set("apiKey", at.apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}
