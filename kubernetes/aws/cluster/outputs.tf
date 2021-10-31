locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_name": "${aws_eks_cluster.platform.id}",
  "${local.component_name_underscores}_security_group_id": "${aws_security_group.eks_cluster.id}",
  "${local.component_name_underscores}_openid_connect_provider_platform_arn": "${aws_iam_openid_connect_provider.platform.arn}",
  "${local.component_name_underscores}_openid_connect_provider_platform_url": "${aws_iam_openid_connect_provider.platform.url}"
}
EOT
}
