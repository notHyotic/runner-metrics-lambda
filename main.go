package main

import (
	"context"
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	// Initialize New Relic agent
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		log.Printf("Processing message ID: %s, Body: %s",
			record.MessageId, record.Body)
		// Add your processing logic here
	}
	return nil
}

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("runner-metrics-lambda"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal("Error creating New Relic application:", err)
	}
	txn := app.StartTransaction("main function")
	defer txn.End()

	lambda.Start(handleRequest)
}
