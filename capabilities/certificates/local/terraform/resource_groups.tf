resource "azurerm_resource_group" "certificates" {
  name     = "certificates-${var.foundation_local_name}"
  location = var.foundation_local_location
}
