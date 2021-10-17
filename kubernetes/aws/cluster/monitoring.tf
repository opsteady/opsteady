resource "aws_cloudwatch_log_group" "platform" {
  # The log group name format is /aws/eks/<cluster-name>/cluster
  # Reference: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
  name              = "/aws/eks/${var.foundation_aws_name}/cluster"
  retention_in_days = 7
}
