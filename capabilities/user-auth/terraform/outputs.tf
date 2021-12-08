locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_oidc_callback_url": "${local.oidc_callback_url}",
  "${local.component_name_underscores}_oidc_url": "${local.oidc_url}",
  "${local.component_name_underscores}_oidc_sp_id": "${azuread_application.oidc.application_id}",
  "${local.component_name_underscores}_oidc_sp_secret": "${azuread_service_principal_password.oidc.value}",
  "${local.component_name_underscores}_primary_domain": "${data.azuread_domains.aad_domains.domains.0.domain_name}"
}
EOT
}
