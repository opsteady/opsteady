resource "azurerm_resource_group" "certificates" {
  name     = "certificates-${var.azure_foundation_name}"
  location = var.azure_foundation_location
}
