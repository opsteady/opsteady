locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_tenant_id": "${var.azuread_tenant_id}",
  "${local.component_name_underscores}_service_principal_id": "${module.certificates.azuread_service_principal_application_id}",
  "${local.component_name_underscores}_service_principal_password": "${module.certificates.azuread_service_principal_password}"
}
EOT
}
