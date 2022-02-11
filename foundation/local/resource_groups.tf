resource "azurerm_resource_group" "foundation" {
  name     = "foundation-${var.local_foundation_name}"
  location = var.local_foundation_location
}
