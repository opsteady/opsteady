resource "aws_security_group" "eks_cluster" {
  name        = "eks-cluster-${var.foundation_aws_name}"
  description = "Cluster communication with worker nodes"
  vpc_id      = var.foundation_aws_vpc_id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "eks_subnets" {
  description       = "Allow the pod subnets to communicate with the EKS API server"
  type              = "ingress"
  protocol          = "tcp"
  from_port         = 443
  to_port           = 443
  security_group_id = aws_security_group.eks_cluster.id

  cidr_blocks = [
    var.foundation_aws_pods_a_cidr_block,
    var.foundation_aws_pods_b_cidr_block,
    var.foundation_aws_pods_c_cidr_block,
  ]
}
