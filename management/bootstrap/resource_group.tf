resource "azurerm_resource_group" "terraform_state" {
  name     = "terraform-state"
  location = var.management_bootstrap_terraform_state_location
}
