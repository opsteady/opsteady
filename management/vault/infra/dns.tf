resource "azurerm_dns_a_record" "vault" {
  name                = "vault"
  zone_name           = data.azurerm_dns_zone.public_management.name
  resource_group_name = "management"
  ttl                 = 300
  records             = [azurerm_public_ip.vault.ip_address]
}

# This resource lives in the nodes-management resource group to ensure that
# the dynamically created loadbalancer can use the IP address.
resource "azurerm_public_ip" "vault" {
  name                = "vault"
  resource_group_name = "nodes-management"
  location            = var.management_infra_location
  allocation_method   = "Static"
  sku                 = "Standard"
}
