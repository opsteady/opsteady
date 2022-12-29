resource "aws_efs_file_system" "platform" {
  creation_token = "efs-platform-${var.aws_foundation_name}"

  encrypted        = true
  performance_mode = "generalPurpose"
  throughput_mode  = "bursting"

  lifecycle_policy {
    transition_to_ia = "AFTER_90_DAYS"
  }

  lifecycle_policy {
    transition_to_primary_storage_class = "AFTER_1_ACCESS"
  }

  tags = {
    "efs.csi.aws.com/cluster" : "true"
  }
}

resource "aws_security_group" "efs" {
  name        = "efs-cluster-${var.aws_foundation_name}"
  description = "EFS communication from worker nodes"
  vpc_id      = var.aws_foundation_vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "nfs" {
  description       = "Allow the pod subnets to communicate with the EFS mount points over NFS"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 2049
  to_port           = 2049
  security_group_id = aws_security_group.efs.id

  cidr_blocks = [
    var.aws_foundation_pods_a_cidr_block,
    var.aws_foundation_pods_b_cidr_block,
    var.aws_foundation_pods_c_cidr_block,
  ]
}

resource "aws_efs_mount_target" "pods_a_subnet" {
  file_system_id  = aws_efs_file_system.platform.id
  subnet_id       = var.aws_foundation_pods_a_subnet_id
  security_groups = [aws_security_group.efs.id]
}

resource "aws_efs_mount_target" "pods_b_subnet" {
  file_system_id  = aws_efs_file_system.platform.id
  subnet_id       = var.aws_foundation_pods_b_subnet_id
  security_groups = [aws_security_group.efs.id]
}

resource "aws_efs_mount_target" "pods_c_subnet" {
  file_system_id  = aws_efs_file_system.platform.id
  subnet_id       = var.aws_foundation_pods_c_subnet_id
  security_groups = [aws_security_group.efs.id]
}
