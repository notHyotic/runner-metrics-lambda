resource "aws_cloudwatch_log_group" "ecs_log_group" {
  name              = "/lambda/runner-metrics-lambda-logs" # This is the log group name you will use in your ECS task definition
  retention_in_days = 7                       # Optional: Set retention (in days) for the logs (e.g., 7 days)
}