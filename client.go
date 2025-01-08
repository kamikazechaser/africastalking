package africastalking

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	req.Header.Set("Content-Type", "application/json")

	return req
}

// postRequestWithCtx builds the HTTP request
func (at *AtClient) postRequestWithCtx(ctx context.Context, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return at.do(req)
}

// do executes the built http request, setting appropriate headers
func (at *AtClient) do(req *http.Request) (*http.Response, error) {
	return at.httpClient.Do(at.setDefaultHeaders(req))
}

// parseResponse is a general utility to decode JSON responses correctly
func parseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("AT server error: code=%s: response_body=%s", resp.Status, string(b))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
