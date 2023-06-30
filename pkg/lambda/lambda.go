package lambda

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/philipep-galdino/liqi-challenge-project/pkg/ethereum"
	"github.com/philipep-galdino/liqi-challenge-project/pkg/sqs"
)

type Transaction struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	for _, message := range sqsEvent.Records {
		var transaction Transaction

		err := json.Unmarshal([]byte(message.Body), &transaction)
		if err != nil {
			log.Printf("Could not unmarshal SQS message body: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		client, err := ethereum.ConnectToEthNetwork()
		if err != nil {
			log.Printf("Failed to connect to Ethereum network: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		privateKeyStr, publicKeyStr, err := ethereum.GenerateKeys()
		if err != nil {
			log.Printf("Failed to generate keys: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		privateKey, err := ethereum.StringToPrivateKey(privateKeyStr)
		if err != nil {
			log.Printf("Failed to parse private key: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		address, err := ethereum.GetAddress(publicKeyStr)
		if err != nil {
			log.Printf("Failed to get address from public key: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		nonce, err := ethereum.GetNonce(client, address)
		if err != nil {
			log.Printf("Failed to get account nonce: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		signedTx, err := ethereum.SignTransaction(transaction.To, transaction.Value, nonce, privateKey)

		if err != nil {
			log.Printf("Failed to create signed transaction: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		hash, err := ethereum.SendTransaction(client, signedTx)
		if err != nil {
			log.Printf("Failed to send transaction: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		messageBody := hash
		messageID, err := sqs.SendMessage(os.Getenv("SQS_QUEUE_URL"), messageBody)
		if err != nil {
			log.Printf("Failed to send SQS message: %v", err)
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: err.Error()}, nil
		}

		log.Println("Transaction hash:", hash)
		log.Println("SQS message ID:", messageID)
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
