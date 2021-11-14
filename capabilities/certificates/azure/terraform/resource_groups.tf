resource "azurerm_resource_group" "certificates" {
  name     = "certificates-${var.foundation_azure_name}"
  location = var.foundation_azure_location
}
