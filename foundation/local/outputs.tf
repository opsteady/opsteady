locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_resource_group": "${azurerm_resource_group.foundation.name}",
  "${local.component_name_underscores}_public_zone_name": "${azurerm_dns_zone.public.name}",
  "${local.component_name_underscores}_public_zone_id": "${azurerm_dns_zone.public.id}"
}
EOT
}
