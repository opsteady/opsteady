# See ADR 0014 Roles responsibilities

# Providing the members is optional as this can also be managed from the Azure Portal
# Providing the owners is optional as this also can be managed from the Azure Portal

resource "azuread_group" "platform_admin" {
  display_name            = "platform-admin"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_platform_admins
  # If no owners are specified, the user who runs the initial Terraform becomes the owner.
  owners = var.management_infra_platform_admin_owners
}

resource "azuread_group" "platform_developer" {
  display_name            = "platform-developer"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_platform_developer_owners
  owners  = var.management_infra_platform_developers
}

resource "azuread_group" "platform_viewer" {
  display_name            = "platform-viewer"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_platform_viewer_owners
  owners  = var.management_infra_platform_viewers
}
