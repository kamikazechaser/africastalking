package africastalking

import (
	"bytes"
	"context"
	"encoding/json"
)

const (
	smsApiPath = "messaging/bulk"
)

type (
	// BulkSMSInput is passed to SendBulkSMS as a parameter.
	BulkSMSInput struct {
		Message      string   `json:"message"`
		SenderID     string   `json:"senderId"`
		PhoneNumbers []string `json:"phoneNumbers"`
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
	var (
		buf             bytes.Buffer
		bulkSMSResponse BulkSMSResponse
	)

	if err := json.NewEncoder(&buf).Encode(struct {
		Username     string   `json:"username"`
		Message      string   `json:"message"`
		SenderID     string   `json:"senderId"`
		PhoneNumbers []string `json:"phoneNumbers"`
	}{
		Username:     at.username,
		Message:      input.Message,
		SenderID:     input.SenderID,
		PhoneNumbers: input.PhoneNumbers,
	}); err != nil {
		return bulkSMSResponse, err
	}

	resp, err := at.postRequestWithCtx(ctx, at.endpoint+smsApiPath, &buf)
	if err != nil {
		return bulkSMSResponse, err
	}

	if err := parseResponse(resp, &bulkSMSResponse); err != nil {
		return bulkSMSResponse, err
	}

	return bulkSMSResponse, nil
}
