data "azurerm_client_config" "current" {}

data "azurerm_key_vault" "management" {
  name                = var.management_infra_key_vault_name
  resource_group_name = "management"
}

data "azurerm_resource_group" "management" {
  name = "management"
}

data "azurerm_kubernetes_cluster" "management" {
  name                = var.management_infra_aks_name
  resource_group_name = "management"
}

data "azurerm_dns_zone" "public_management" {
  name                = "management.${var.management_infra_domain}"
  resource_group_name = "management"
}
