package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type WorkflowBody struct {
	RunID        string `json:"run_id"`
	Repository   string `json:"repository"`
	RunStartedAt string `json:"run_started_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Conclusion   string `json:"conclusion"`
}

func handleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		log.Println("Processing message")
		log.Printf("Message ID: %s", record.MessageId)
		log.Printf("Message Body: %s", record.Body)

		var body WorkflowBody
		if err := json.Unmarshal([]byte(record.Body), &body); err != nil {
			log.Printf("Failed to unmarshal message body: %v", err)
			continue
		}

		startedAt, err1 := time.Parse(time.RFC3339, body.RunStartedAt)
		updatedAt, err2 := time.Parse(time.RFC3339, body.UpdatedAt)
		if err1 != nil || err2 != nil {
			log.Printf("Failed to parse timestamps: %v, %v", err1, err2)
			continue
		}

		duration := updatedAt.Sub(startedAt)
		log.Printf("Workflow run time: %s", duration)

 		// Send custom metric to CloudWatch
        sess := session.Must(session.NewSession())
        cw := cloudwatch.New(sess)
        _, err := cw.PutMetricData(&cloudwatch.PutMetricDataInput{
            Namespace: aws.String("RunnerMetrics"),
            MetricData: []*cloudwatch.MetricDatum{
                {
                    MetricName: aws.String("WorkflowDurationSeconds"),
                    Value:      aws.Float64(float64(duration.Seconds())),
                    Unit:       aws.String("Seconds"),
                    Dimensions: []*cloudwatch.Dimension{
                        {
                            Name:  aws.String("Repository"),
                            Value: aws.String(body.Repository),
                        },
                        {
                            Name:  aws.String("Conclusion"),
                            Value: aws.String(body.Conclusion),
                        },
                    },
                },
            },
        })
        if err != nil {
            log.Printf("Failed to put custom metric: %v", err)
        }

		// Extract and log job status
        switch body.Conclusion {
        case "success":
            log.Printf("Job status: SUCCESS")
        case "cancelled", "canceled":
            log.Printf("Job status: CANCELLED")
        default:
            log.Printf("Job status: FAILURE (%s)", body.Conclusion)
        }
	}
	return nil
}

func main() {
	lambda.Start(handleRequest)
}
