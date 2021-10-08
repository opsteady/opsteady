data "azurerm_dns_zone" "public_root" {
  provider = azurerm.management

  name                = var.management_infra_domain
  resource_group_name = "management"
}

data "azurerm_client_config" "current" {}
