resource "azurerm_resource_group" "kubernetes" {
  name     = "kubernetes-${var.foundation_azure_name}"
  location = var.foundation_azure_location
}
