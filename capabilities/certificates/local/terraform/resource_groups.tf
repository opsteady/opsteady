resource "azurerm_resource_group" "certificates" {
  name     = "certificates-${var.local_foundation_name}"
  location = var.local_foundation_location
}
