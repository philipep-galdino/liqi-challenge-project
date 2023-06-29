package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Transaction struct {
	To    string `json:"to"`
	Value string `json: "value"`
}

func SendMessage(queueURL string, messageBody string) (*sqs.SendMessageOutput, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)

	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(messageBody),
		QueueUrl:     aws.String(queueURL),
	})

	return result, err
}
