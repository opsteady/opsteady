resource "azurerm_resource_group" "dns" {
  name     = "dns-${var.azure_foundation_name}"
  location = var.azure_foundation_location
}
