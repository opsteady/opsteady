resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_iam_role_arn": "${aws_iam_role.aws_ebs_csi_driver.arn}"
}
EOT
}
