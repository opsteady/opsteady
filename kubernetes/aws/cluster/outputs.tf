resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_name": "${aws_eks_cluster.platform.id}",
  "${var.platform_vault_vars_name}_security_group_id": "${aws_security_group.eks_cluster.id}",
  "${var.platform_vault_vars_name}_openid_connect_provider_platform_arn": "${aws_iam_openid_connect_provider.platform.arn}",
  "${var.platform_vault_vars_name}_openid_connect_provider_platform_url": "${aws_iam_openid_connect_provider.platform.url}"
}
EOT
}
