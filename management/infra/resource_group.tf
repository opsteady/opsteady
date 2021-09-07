resource "azurerm_resource_group" "management" {
  name     = "management"
  location = var.management_infra_location
}
