resource "aws_eks_addon" "vpc_cni" {
  addon_name               = "vpc-cni"
  addon_version            = "v1.9.1-eksbuild.1"
  resolve_conflicts        = "OVERWRITE"
  cluster_name             = aws_eks_cluster.platform.name
  service_account_role_arn = aws_iam_role.eks_cni.arn
}

resource "aws_eks_addon" "kube_proxy" {
  addon_name               = "kube-proxy"
  addon_version            = "v1.21.2-eksbuild.2"
  resolve_conflicts        = "OVERWRITE"
  cluster_name             = aws_eks_cluster.platform.name
}

resource "aws_eks_addon" "coredns" {
  addon_name               = "coredns"
  addon_version            = "v1.8.4-eksbuild.1"
  resolve_conflicts        = "OVERWRITE"
  cluster_name             = aws_eks_cluster.platform.name

  depends_on = [
    aws_eks_node_group.system
  ]
}
