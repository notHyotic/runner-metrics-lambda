terraform {
  backend "s3" {
    bucket         = "shared-terraform-state-bucket1" # Replace with your S3 bucket name
    key            = "state/runner-metrics-lambda.tfstate" # The path within the bucket to store the state file
    region         = "us-east-1"                      # Your AWS region
    encrypt        = true                             # Encrypt the state file in S3
    dynamodb_table = "terraform-state-lock-table"     # DynamoDB table for state locking
    acl            = "private"                        # ACL for the S3 bucket (private is recommended)
  }
}