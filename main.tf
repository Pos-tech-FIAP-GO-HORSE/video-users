provider "aws" {
  region = var.aws_region
}

resource "aws_sns_topic" "video_user_topic" {
  name = "video-user"
}

resource "aws_iam_role" "lambda_exec" {
  name = "video_user_lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = { Service = "lambda.amazonaws.com" },
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_policy" "lambda_policy" {
  name = "video_user_access_policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = ["sns:Publish"],
        Resource = [
          aws_sns_topic.video_user_topic.arn
        ]
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_custom_policy" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}

resource "aws_lambda_function" "video_user" {
  function_name = "video-user"

  role    = aws_iam_role.lambda_exec.arn
  handler = "main"
  runtime = "provided.al2023"
  timeout = 300

  filename         = "build/video-users.zip"
  source_code_hash = filebase64sha256("build/video-users.zip")
}
