# runner metrics lambda

## useful query

```
SELECT average(duration_ms) FROM WorkflowRun
```

```
SELECT percentage(count(*), WHERE conclusion = 'success') AS 'Success Rate' FROM WorkflowRun
```

```
SELECT count(*) FROM WorkflowRun
```

## Test event
```
{
  "Records": [
    {
      "messageId": "1",
      "receiptHandle": "abc",
      "body": "{\"run_id\": \"1234567890\", \"repository\": \"your-org/your-repo\", \"run_started_at\": \"2025-05-15T12:34:56Z\", \"updated_at\": \"2025-05-15T12:45:00Z\", \"status\": \"completed\", \"conclusion\": \"success\"}",
      "attributes": {},
      "messageAttributes": {},
      "md5OfBody": "",
      "eventSource": "aws:sqs",
      "eventSourceARN": "arn:aws:sqs:us-east-1:123456789012:queue-name",
      "awsRegion": "us-east-1"
    }
  ]
}
```