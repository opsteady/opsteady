# DNS used for management records
resource "azurerm_dns_zone" "public_root" {
  name                = var.management_infra_domain
  resource_group_name = azurerm_resource_group.management.name
}

resource "azurerm_dns_zone" "public_management" {
  name                = "management.${var.management_infra_domain}"
  resource_group_name = azurerm_resource_group.management.name
}

resource "azurerm_dns_ns_record" "public_management_delegation" {
  name                = "management"
  zone_name           = azurerm_dns_zone.public_root.name
  resource_group_name = azurerm_resource_group.management.name
  ttl                 = 300

  records = azurerm_dns_zone.public_management.name_servers
}
