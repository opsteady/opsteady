resource "aws_iam_role" "aws_efs_csi_driver" {
  name = "aws-efs-csi-driver-${var.foundation_aws_name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "${var.kubernetes_aws_cluster_openid_connect_provider_platform_arn}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${replace(var.kubernetes_aws_cluster_openid_connect_provider_platform_url, "https://", "")}:sub": ["system:serviceaccount:platform:efs-csi-controller-sa", "system:serviceaccount:platform:efs-csi-node-sa"]
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_policy" "aws_efs_csi_driver" {
  name        = "aws-efs-csi-driver-${var.foundation_aws_name}"
  path        = "/"
  description = "AWS EFS CSI driver"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "elasticfilesystem:DescribeAccessPoints",
        "elasticfilesystem:DescribeFileSystems",
        "elasticfilesystem:DescribeMountTargets",
        "ec2:DescribeAvailabilityZones"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "elasticfilesystem:CreateAccessPoint"
      ],
      "Resource": "*",
      "Condition": {
        "StringLike": {
          "aws:RequestTag/efs.csi.aws.com/cluster": "true"
        }
      }
    },
    {
      "Effect": "Allow",
      "Action": "elasticfilesystem:DeleteAccessPoint",
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "aws:ResourceTag/efs.csi.aws.com/cluster": "true"
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "aws_efs_csi_driver" {
  name        = "aws-efs-csi-driver-${var.foundation_aws_name}"
  roles      = [aws_iam_role.aws_efs_csi_driver.name]
  policy_arn = aws_iam_policy.aws_efs_csi_driver.arn
}
