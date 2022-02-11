resource "azurerm_resource_group" "dns" {
  name     = "dns-${var.local_foundation_name}"
  location = var.local_foundation_location
}
