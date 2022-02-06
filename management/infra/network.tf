# VNET for management AKS
resource "azurerm_virtual_network" "management" {
  name                = "management"
  location            = azurerm_resource_group.management.location
  resource_group_name = azurerm_resource_group.management.name
  address_space       = var.management_infra_vnet_address_space
}

resource "azurerm_subnet" "pods" {
  name                 = "pods-management"
  resource_group_name  = azurerm_resource_group.management.name
  virtual_network_name = azurerm_virtual_network.management.name
  address_prefixes     = var.management_infra_azure_subnet_pods_address_prefixes

  enforce_private_link_endpoint_network_policies = true
  service_endpoints = [
    "Microsoft.ContainerRegistry",
    "Microsoft.KeyVault",
    "Microsoft.Storage"
  ]
}

# Although we are not using this yet, we've decided to create it if we want to use it later
resource "azurerm_subnet" "pub" {
  name                 = "pub-management"
  resource_group_name  = azurerm_resource_group.management.name
  virtual_network_name = azurerm_virtual_network.management.name
  address_prefixes     = var.management_infra_azure_subnet_public_address_prefixes
}

resource "azurerm_network_security_group" "pub" {
  name                = "pub-management"
  location            = azurerm_resource_group.management.location
  resource_group_name = azurerm_resource_group.management.name
}

resource "azurerm_subnet_network_security_group_association" "pub" {
  subnet_id                 = azurerm_subnet.pub.id
  network_security_group_id = azurerm_network_security_group.pub.id
}

# Reserve some IP space for outgoing traffic on the AKS loadbalancer
resource "azurerm_public_ip_prefix" "pub" {
  count = var.management_infra_minimal ? 0 : 1

  name                = "pub-${var.management_infra_aks_name}"
  location            = azurerm_resource_group.management.location
  resource_group_name = azurerm_resource_group.management.name

  prefix_length = 29
}
