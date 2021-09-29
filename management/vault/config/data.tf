data "azurerm_client_config" "current" {}

data "azuread_application_published_app_ids" "well_known" {}

data "azuread_group" "platform_admin" {
  display_name = "platform-admin"
}

data "azuread_group" "platform_operator" {
  display_name = "platform-operator"
}

data "azuread_group" "platform_viewer" {
  display_name = "platform-viewer"
}
