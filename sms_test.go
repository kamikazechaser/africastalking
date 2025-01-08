package africastalking

import (
	"context"
	"os"
	"testing"
)

// E2E Test
func TestAfricasTalking_SendBulkSMS(t *testing.T) {
	apiKey := os.Getenv("TEST_AT_API_KEY")
	username := os.Getenv("TEST_AT_USERNAME")
	sandbox := false

	testAtClient := New(apiKey, username, sandbox)

	testData := BulkSMSInput{
		PhoneNumbers: []string{
			"+254723123123",
		},
		SenderID: "",
		Message:  "kamikazechaser/africastalking test",
	}

	resp, err := testAtClient.SendBulkSMS(context.Background(), testData)
	if err != nil {
		t.Fatalf("Failed: %+v\n", err)
	}

	if resp.SMSMessageData.Recipients[0].Status != "Success" {
		t.Fatalf("Failed: %+v\n", err)
	}

	t.Logf("Success: %+v\n", resp)
}
