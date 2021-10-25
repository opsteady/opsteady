locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_iam_role_arn": "${aws_iam_role.aws_efs_csi_driver.arn}",
  "${local.component_name_underscores}_filesystem_id": "${aws_efs_file_system.platform.id}"
}
EOT
}
