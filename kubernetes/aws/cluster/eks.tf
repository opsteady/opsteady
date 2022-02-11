resource "aws_kms_key" "platform" {
  description         = "eks-platform-${var.aws_foundation_name}"
  enable_key_rotation = true
}

resource "aws_eks_cluster" "platform" {
  name     = var.aws_foundation_name
  role_arn = aws_iam_role.eks.arn
  version  = var.aws_cluster_kubernetes_version

  enabled_cluster_log_types = ["api", "audit", "controllerManager"]

  encryption_config {
    provider {
      key_arn = aws_kms_key.platform.arn
    }
    resources = ["secrets"]
  }

  vpc_config {
    // By enabling this endpoint, resources in private networks (pods, nodes, etc.) can talk
    // directly to the Kubernetes API. The public endpoint is also enabled but this can then
    // be used for management purposes only and, optionally, controlled via CIDR restrictions.
    endpoint_private_access = true
    public_access_cidrs     = var.aws_cluster_public_access_cidrs
    security_group_ids      = [aws_security_group.eks_cluster.id]
    subnet_ids = [
      var.aws_foundation_eks_a_subnet_id,
      var.aws_foundation_eks_b_subnet_id,
      var.aws_foundation_eks_c_subnet_id,
    ]
  }

  kubernetes_network_config {
    // In the future this can potentially be adjusted if it clashes with peered networks.
    service_ipv4_cidr = var.aws_cluster_service_ipv4_cidr
  }

  depends_on = [
    aws_iam_role_policy_attachment.eks_cluster,
    aws_iam_role_policy_attachment.eks_service,
    aws_cloudwatch_log_group.platform,
  ]
}

resource "aws_eks_node_group" "system" {
  cluster_name    = var.aws_foundation_name
  node_group_name = "system"
  instance_types  = var.aws_cluster_system_node_group_instance_types
  node_role_arn   = aws_iam_role.eks_system_node_group.arn
  labels = {
    name = "system"
  }

  subnet_ids = [
    var.aws_foundation_pods_a_subnet_id,
    var.aws_foundation_pods_b_subnet_id,
    var.aws_foundation_pods_c_subnet_id,
  ]

  launch_template {
    id      = aws_launch_template.system.id
    version = aws_launch_template.system.latest_version
  }

  scaling_config {
    desired_size = var.aws_cluster_system_node_group_node_count
    min_size     = var.aws_cluster_system_node_group_node_count
    max_size     = var.aws_cluster_system_node_group_node_count
  }

  update_config {
    // We want minimal disruption so we accept that it will take
    // a bit longer to update the nodes.
    max_unavailable = 1
  }

  # Ensure that IAM Role permissions are created before and deleted after EKS Node Group handling.
  # Otherwise, EKS will not be able to properly delete EC2 Instances and Elastic Network Interfaces.
  depends_on = [
    aws_eks_cluster.platform,
    aws_iam_role_policy_attachment.eks_worker_node,
    aws_iam_role_policy_attachment.eks_cni,
    aws_iam_role_policy_attachment.eks_container_registry,
  ]
}

# -------------
# NOTE: This is disabled because Terraform currently can't reconcile a KMS key with multiple policies:
# https://github.com/hashicorp/terraform-provider-aws/issues/21225
#
# Re-enable when this is fixed: https://github.com/opsteady/opsteady/issues/65
# -------------
#
# resource "aws_kms_key" "system" {
#   description         = "eks-system-${var.aws_foundation_name}"
#   enable_key_rotation = true
#   policy = <<POLICY
# {
#   "Version": "2012-10-17",
#   "Statement": [
#     {
#       "Effect": "Allow",
#       "Principal": {
#         "AWS": [
#           "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
#         ]
#       },
#       "Action": [
#         "kms:*"
#       ],
#       "Resource": "*"
#     },
#     {
#       "Effect": "Allow",
#       "Principal": {
#         "AWS": [
#           "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
#         ]
#       },
#       "Action": [
#         "kms:Encrypt",
#         "kms:Decrypt",
#         "kms:ReEncrypt*",
#         "kms:GenerateDataKey*",
#         "kms:DescribeKey"
#       ],
#       "Resource": "*"
#     },
#     {
#       "Effect": "Allow",
#       "Principal": {
#         "AWS": [
#           "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/aws-service-role/autoscaling.amazonaws.com/AWSServiceRoleForAutoScaling"
#         ]
#       },
#       "Action": [
#         "kms:CreateGrant"
#       ],
#       "Resource": "*",
#       "Condition": {
#         "Bool": {
#           "kms:GrantIsForAWSResource": true
#         }
#       }
#     }
#   ]
# }
# POLICY
# }

resource "aws_launch_template" "system" {
  image_id               = data.aws_ssm_parameter.eks_ami_id.value
  name                   = "${var.aws_foundation_name}-system"
  update_default_version = true

  block_device_mappings {
    device_name = "/dev/xvda"

    ebs {
      encrypted = false
      # kms_key_id  = aws_kms_key.system.arn
      volume_type = "gp2"
      volume_size = 20
    }
  }

  user_data = base64encode(templatefile("template/userdata.tpl", {
    CLUSTER_NAME   = aws_eks_cluster.platform.name,
    B64_CLUSTER_CA = aws_eks_cluster.platform.certificate_authority[0].data,
    API_SERVER_URL = aws_eks_cluster.platform.endpoint
  }))

  lifecycle {
    create_before_destroy = true
  }
}
