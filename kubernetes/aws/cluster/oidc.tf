data "tls_certificate" "eks_oidc_issuer" {
  url = aws_eks_cluster.platform.identity.0.oidc.0.issuer
}

resource "aws_iam_openid_connect_provider" "platform" {
  client_id_list  = ["sts.${data.aws_partition.current.dns_suffix}"]
  thumbprint_list = [data.tls_certificate.eks_oidc_issuer.certificates.0.sha1_fingerprint]
  url             = aws_eks_cluster.platform.identity.0.oidc.0.issuer
}
