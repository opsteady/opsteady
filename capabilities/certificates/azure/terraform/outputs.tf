resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_msi_id": "${azurerm_user_assigned_identity.certificates.id}",
  "${var.platform_vault_vars_name}_msi_client_id": "${azurerm_user_assigned_identity.certificates.client_id}"
}
EOT
}
