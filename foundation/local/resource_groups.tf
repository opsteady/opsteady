resource "azurerm_resource_group" "foundation" {
  name     = "foundation-${var.foundation_local_name}"
  location = var.foundation_local_location
}
