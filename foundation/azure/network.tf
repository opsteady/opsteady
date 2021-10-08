resource "azurerm_virtual_network" "platform" {
  name                = var.foundation_azure_name
  location            = azurerm_resource_group.foundation.location
  resource_group_name = azurerm_resource_group.foundation.name
  address_space       = var.foundation_azure_vnet_address_space
}

resource "azurerm_subnet" "pods" {
  name                 = "pods-${var.foundation_azure_name}"
  resource_group_name  = azurerm_resource_group.foundation.name
  virtual_network_name = azurerm_virtual_network.platform.name
  address_prefixes     = var.foundation_azure_subnet_pods_address_prefixes

  enforce_private_link_endpoint_network_policies = true
  service_endpoints = [
    "Microsoft.ContainerRegistry",
    "Microsoft.KeyVault",
    "Microsoft.Storage"
  ]
}

resource "azurerm_subnet" "pub" {
  name                 = "pub-${var.foundation_azure_name}"
  resource_group_name  = azurerm_resource_group.foundation.name
  virtual_network_name = azurerm_virtual_network.platform.name
  address_prefixes     = var.foundation_azure_subnet_public_address_prefixes
}

# Although we are not using this yet, we've decided to create it if we want to use it later
resource "azurerm_network_security_group" "pub" {
  name                = "pub-${var.foundation_azure_name}"
  location            = azurerm_resource_group.foundation.location
  resource_group_name = azurerm_resource_group.foundation.name
}

resource "azurerm_subnet_network_security_group_association" "pub" {
  subnet_id                 = azurerm_subnet.pub.id
  network_security_group_id = azurerm_network_security_group.pub.id
}

# Reserve some IP space for outgoing traffic on the AKS loadbalancer
resource "azurerm_public_ip_prefix" "pub" {
  name                = "pub-${var.foundation_azure_name}"
  location            = azurerm_resource_group.foundation.location
  resource_group_name = azurerm_resource_group.foundation.name

  prefix_length = 29
}
