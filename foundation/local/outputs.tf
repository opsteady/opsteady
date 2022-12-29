resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_resource_group": "${azurerm_resource_group.foundation.name}",
  "${var.platform_vault_vars_name}_public_zone_name": "${azurerm_dns_zone.public.name}",
  "${var.platform_vault_vars_name}_public_zone_id": "${azurerm_dns_zone.public.id}"
}
EOT
}
