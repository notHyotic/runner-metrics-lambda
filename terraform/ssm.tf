resource "aws_ssm_parameter" "new_relic_license_key" {
  name        = "/runner-metrics-lambda/newrelic/license_key"
  type        = "SecureString"
  value       = "unintialized"
  description = "New Relic License Key for runner-metrics-lambda"
  lifecycle {
    ignore_changes = [value]
  }
}

resource "aws_ssm_parameter" "new_relic_user_key" {
  name        = "/runner-metrics-lambda/newrelic/user_key"
  type        = "SecureString"
  value       = "unintialized"
  description = "New Relic User Key for runner-metrics-lambda"
  lifecycle {
    ignore_changes = [value]
  }
}