# Create a unique subdomain for this platform
resource "azurerm_dns_zone" "public" {
  name                = "${var.foundation_azure_name}.${var.management_infra_domain}"
  resource_group_name = azurerm_resource_group.foundation.name
}

resource "azurerm_dns_ns_record" "public_management_delegation" {
  provider = azurerm.management

  name                = var.foundation_azure_name
  zone_name           = data.azurerm_dns_zone.public_root.name
  resource_group_name = "management"
  ttl                 = 300

  records = azurerm_dns_zone.public.name_servers
}
