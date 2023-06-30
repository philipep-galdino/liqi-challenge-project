package lambda

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	request := events.SQSEvent{
		Records: []events.SQSMessage{
			{
				Body: `{"to":"0xAbC123...", "value":"0x1bc16d674ec80000"}`,
			},
		},
	}

	_, err := HandleRequest(context.Background(), request)
	if err != nil {
		t.Errorf("HandleRequest failed: %v", err)
	}
}
