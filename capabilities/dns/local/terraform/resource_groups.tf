resource "azurerm_resource_group" "dns" {
  name     = "dns-${var.foundation_local_name}"
  location = var.foundation_local_location
}
