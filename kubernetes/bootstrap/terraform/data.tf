data "azuread_group" "platform_admin" {
  display_name = "platform-admin"
  security_enabled = true
}

data "azuread_group" "platform_operator" {
  display_name = "platform-operator"
  security_enabled = true
}

data "azuread_group" "platform_viewer" {
  display_name = "platform-viewer"
  security_enabled = true
}
