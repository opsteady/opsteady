resource "azurerm_resource_group" "dns" {
  name     = "dns-${var.foundation_azure_name}"
  location = var.foundation_azure_location
}
