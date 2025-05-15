package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var app *newrelic.Application

type WorkflowBody struct {
	RunID        string `json:"run_id"`
	Repository   string `json:"repository"`
	RunStartedAt string `json:"run_started_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Conclusion   string `json:"conclusion"`
}

func init() {
	var err error
	app, err = newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEWRELIC_APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		log.Fatalf("failed to create New Relic app: %v", err)
	}
	log.Printf("App Name: %s", os.Getenv("NEWRELIC_APP_NAME"))
	log.Printf("New Relic app initialized")

	if err := app.WaitForConnection(5 * time.Second); err != nil {
		log.Printf("New Relic app failed to connect: %v", err)
	}
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	txn := app.StartTransaction("SQS_Lambda_Handler")
	defer txn.End()

	for _, record := range sqsEvent.Records {
		seg := txn.StartSegment("Process SQS Message")
		log.Println("Processing message")
		log.Printf("Message ID: %s", record.MessageId)
		log.Printf("Message Body: %s", record.Body)

		var body WorkflowBody
		if err := json.Unmarshal([]byte(record.Body), &body); err != nil {
			log.Printf("Failed to unmarshal message body: %v", err)
			seg.End()
			continue
		}

		startedAt, err1 := time.Parse(time.RFC3339, body.RunStartedAt)
		updatedAt, err2 := time.Parse(time.RFC3339, body.UpdatedAt)
		if err1 != nil || err2 != nil {
			log.Printf("Failed to parse timestamps: %v, %v", err1, err2)
			seg.End()
			continue
		}

		duration := updatedAt.Sub(startedAt)
		log.Printf("Workflow run time: %s", duration)

		seg.End()
	}
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
