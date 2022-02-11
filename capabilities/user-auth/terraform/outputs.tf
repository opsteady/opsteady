resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_oidc_callback_url": "${local.oidc_callback_url}",
  "${var.platform_vault_vars_name}_oidc_url": "${local.oidc_url}",
  "${var.platform_vault_vars_name}_oidc_sp_id": "${azuread_application.oidc.application_id}",
  "${var.platform_vault_vars_name}_oidc_sp_secret": "${azuread_service_principal_password.oidc.value}",
  "${var.platform_vault_vars_name}_primary_domain": "${data.azuread_domains.aad_domains.domains.0.domain_name}"
}
EOT
}
