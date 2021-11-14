resource "aws_iam_role" "certificates" {
  name = "eks-certificates-${var.foundation_aws_name}"

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
          "${replace(var.kubernetes_aws_cluster_openid_connect_provider_platform_url, "https://", "")}:sub": "system:serviceaccount:platform:cert-manager"
        }
      }
    }
  ]
}
EOF
}

data "aws_iam_policy_document" "certificates_policy" {
  statement {
    actions = [
      "route53:ChangeResourceRecordSets",
      "route53:ListResourceRecordSets"
    ]

    resources = [
      "arn:aws:route53:::hostedzone/${var.foundation_aws_public_zone_id}",
    ]
  }

  statement {
    actions = [
      "route53:GetChange"
    ]

    resources = [
      "arn:aws:route53:::change/*"
    ]
  }

  statement {
    actions = [
      "route53:ListHostedZonesByName"
    ]

    resources = [
      "*"
    ]
  }
}

resource "aws_iam_policy" "certificates" {
  name   = "certificates_policy"
  path   = "/"
  policy = data.aws_iam_policy_document.certificates_policy.json
}

resource "aws_iam_role_policy_attachment" "certificates" {
  role       = aws_iam_role.certificates.name
  policy_arn = aws_iam_policy.certificates.arn
}
