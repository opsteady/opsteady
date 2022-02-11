resource "aws_iam_role" "external_dns" {
  name = "external-dns-${var.aws_foundation_name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "${var.aws_cluster_openid_connect_provider_platform_arn}"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
        "StringEquals": {
          "${replace(var.aws_cluster_openid_connect_provider_platform_url, "https://", "")}:sub": "system:serviceaccount:platform:external-dns"
        }
      }
    }
  ]
}
EOF
}

resource "aws_iam_policy" "external_dns" {
  name        = "external-dns-${var.aws_foundation_name}"
  path        = "/"
  description = "External DNS"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "route53:ChangeResourceRecordSets"
      ],
      "Resource": [
        "arn:aws:route53:::hostedzone/${var.aws_foundation_public_zone_id}"
      ]
    },
    {
      "Effect": "Allow",
      "Action": [
        "route53:ListHostedZones",
        "route53:ListResourceRecordSets"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "aws_load_balancer_controller" {
  name       = "external-dns-${var.aws_foundation_name}"
  roles      = [aws_iam_role.external_dns.name]
  policy_arn = aws_iam_policy.external_dns.arn
}
