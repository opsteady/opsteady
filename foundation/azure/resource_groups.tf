# TODO: The name needs to be outputed to Vault as others are using it as well
resource "azurerm_resource_group" "foundation" {
  name     = "foundation-${var.foundation_azure_name}"
  location = var.foundation_azure_location
}
