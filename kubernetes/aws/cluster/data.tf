data "aws_partition" "current" {}

data "aws_caller_identity" "current" {}

data "aws_ssm_parameter" "eks_ami_id" {
  name = "/aws/service/eks/optimized-ami/${aws_eks_cluster.platform.version}/amazon-linux-2/recommended/image_id"
}
