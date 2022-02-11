resource "azurerm_resource_group" "foundation" {
  name     = "foundation-${var.azure_foundation_name}"
  location = var.azure_foundation_location
}
