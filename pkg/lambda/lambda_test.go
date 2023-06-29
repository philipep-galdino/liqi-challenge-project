package lambda

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	request := events.SQSEvent{}

	HandleRequest(request)
}
