resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_tenant_id": "${var.azuread_tenant_id}",
  "${var.platform_vault_vars_name}_service_principal_id": "${module.certificates.azuread_service_principal_application_id}",
  "${var.platform_vault_vars_name}_service_principal_password": "${module.certificates.azuread_service_principal_password}"
}
EOT
}
