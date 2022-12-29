resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_resource_group": "${azurerm_resource_group.foundation.name}",
  "${var.platform_vault_vars_name}_pods_subnet_id": "${azurerm_subnet.pods.id}",
  "${var.platform_vault_vars_name}_key_vault_id": "${azurerm_key_vault.platform.id}",
  "${var.platform_vault_vars_name}_public_ip_prefix_id": "${azurerm_public_ip_prefix.pub.id}",
  "${var.platform_vault_vars_name}_log_analytics_id": "${azurerm_log_analytics_workspace.platform.id}",
  "${var.platform_vault_vars_name}_public_zone_name": "${azurerm_dns_zone.public.name}",
  "${var.platform_vault_vars_name}_public_zone_id": "${azurerm_dns_zone.public.id}"
}
EOT
}
