locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_msi_id": "${azurerm_user_assigned_identity.certificates.id}",
  "${local.component_name_underscores}_msi_client_id": "${azurerm_user_assigned_identity.certificates.client_id}"
}
EOT
}
