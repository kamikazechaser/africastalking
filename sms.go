package africastalking

import (
	"bytes"
	"context"
	"net/url"
	"strings"
)

const (
	smsApiPath = "messaging"
)

type (
	// BulkSMSInput is passed to SendBulkSMS as a parameter.
	BulkSMSInput struct {
		To      []string
		Message string
		From    string
	}

	// BulkSMSRecipient is returned as part of the BulkSMSResponse.
	BulkSMSRecipient struct {
		StatusCode uint   `json:"statusCode"`
		Number     string `json:"number"`
		Status     string `json:"status"`
		Cost       string `json:"cost"`
		MessageID  string `json:"messageId"`
	}

	// BulkSMSResponse is returned by SendBulkSMS as a response.
	BulkSMSResponse struct {
		SMSMessageData struct {
			Message    string             `json:"Message"`
			Recipients []BulkSMSRecipient `json:"Recipients"`
		} `json:"SMSMessageData"`
	}
)

// SendBulkSMS makes a POST request to send bulk SMS's the Africa's Talking and returns a response.
// It uses opinionated defaults.
func (at *AtClient) SendBulkSMS(ctx context.Context, input BulkSMSInput) (BulkSMSResponse, error) {
	bulkSMSResponse := BulkSMSResponse{}

	form := url.Values{
		"username":             {at.username},
		"to":                   {strings.Join(input.To, ",")},
		"message":              {input.Message},
		"bulkSMSMode":          {"1"},
		"enqueue":              {"1"},
		"retryDurationInHours": {"1"},
	}

	if input.From != "" {
		form.Set("from", input.From)
	}

	resp, err := at.postRequestWithCtx(ctx, at.endpoint+smsApiPath, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return bulkSMSResponse, err
	}

	if err := parseResponse(resp, &bulkSMSResponse); err != nil {
		return bulkSMSResponse, err
	}

	return bulkSMSResponse, nil
}
