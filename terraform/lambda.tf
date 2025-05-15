data "aws_ssm_parameter" "new_relic_license_key" {
  name            = aws_ssm_parameter.new_relic_license_key.name
  with_decryption = true
}

data "aws_ssm_parameter" "new_relic_user_key" {
  name            = aws_ssm_parameter.new_relic_user_key.name
  with_decryption = true
}

resource "aws_lambda_function" "runner_metrics" {
  function_name = "runner-metrics-lambda"
  role          = aws_iam_role.lambda_exec.arn
  runtime = "provided.al2023"
  handler = "bootstrap"

  # Update these to match your deployment package location
  filename         = "../build/runner-metrics-lambda.zip"
  source_code_hash = filebase64sha256("../build/runner-metrics-lambda.zip")

  environment {
    variables = {
      NEWRELIC_LICENSE_KEY = data.aws_ssm_parameter.new_relic_license_key.value
      NEWRELIC_USER_KEY    = data.aws_ssm_parameter.new_relic_user_key.value
    }
  }
}

resource "aws_lambda_event_source_mapping" "sqs_event" {
  event_source_arn = aws_sqs_queue.runner_metrics_queue.arn
  function_name    = aws_lambda_function.runner_metrics.arn
  batch_size       = 10
  enabled          = true
}