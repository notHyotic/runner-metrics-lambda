resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# resource "aws_iam_role_policy_attachment" "lambda_sqs" {
#   role       = aws_iam_role.lambda_exec.name
#   policy_arn = "arn:aws:iam::aws:policy/AWSLambdaSQSQueueExecutionRole"
# }

resource "aws_iam_role_policy" "lambda_cloudwatch_put_metric" {
  name = "lambda-cloudwatch-put-metric"
  role = aws_iam_role.lambda_exec.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "cloudwatch:PutMetricData"
        ],
        Resource = "*"
      }
    ]
  })
}