package africastalking

import (
	"context"
	"os"
	"testing"
)

func TestAfricasTalking_SendBulkSMS(t *testing.T) {
	apiKey := os.Getenv("TEST_AT_API_KEY")
	username := os.Getenv("TEST_AT_USERNAME")
	sandbox := true

	testAtClient := New(apiKey, username, sandbox)

	testData := BulkSMSInput{
		To: []string{
			"+254722123123",
			"+254723123123",
		},
		From:    "",
		Message: "kamikazechaser/africastalking test",
	}

	resp, err := testAtClient.SendBulkSMS(context.Background(), testData)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}

	if resp.SMSMessageData.Recipients[0].Status != "Success" {
		t.Fatalf("Failed: %v", err)
	}
}
