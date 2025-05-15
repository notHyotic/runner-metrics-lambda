resource "aws_sqs_queue" "runner_metrics_queue" {
  name = "runner-metrics-queue"
}