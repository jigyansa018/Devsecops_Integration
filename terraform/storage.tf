resource "aws_s3_bucket" "tf_state" {
  bucket = "chat-app-terraform-state"
  force_destroy = true
}

resource "aws_dynamodb_table" "tf_lock" {
  name         = "chat-app-terraform-lock"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}
