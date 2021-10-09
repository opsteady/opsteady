resource "azurerm_resource_group" "foundation" {
  name     = "foundation-${var.foundation_azure_name}"
  location = var.foundation_azure_location
}
