resource "azurerm_resource_group" "kubernetes" {
  name     = "kubernetes-${var.azure_foundation_name}"
  location = var.azure_foundation_location
}
