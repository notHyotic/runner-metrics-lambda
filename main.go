package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		log.Printf("Processing message ID: %s, Body: %s",
			record.MessageId, record.Body)
		// Add your processing logic here
	}
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
