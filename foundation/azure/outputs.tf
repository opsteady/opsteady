locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_resource_group": "${azurerm_resource_group.foundation.name}",
  "${local.component_name_underscores}_pods_subnet_id": "${azurerm_subnet.pods.id}",
  "${local.component_name_underscores}_key_vault_id": "${azurerm_key_vault.platform.id}",
  "${local.component_name_underscores}_public_ip_prefix_id": "${azurerm_public_ip_prefix.pub.id}",
  "${local.component_name_underscores}_log_analytics_id": "${azurerm_log_analytics_workspace.platform.id}"
}
EOT
}
