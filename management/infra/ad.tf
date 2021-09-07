# See ADR 0014 Roles responisbilites

# Providing the members is optional as this can also be managed from the Azure Portal
# Providing the owners is optional as this also can be managed from the Azure Portal


resource "azuread_group" "cluster_admin" {
  display_name            = "cluster-admin"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_cluster_admins
  # If no owners specified ther user who runs the initial Terraform becomes the owner
  owners = var.management_infra_cluster_admin_owners
}

resource "azuread_group" "cluster_developer" {
  display_name            = "cluster-developer"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_cluster_developer_owners
  owners  = var.management_infra_cluster_developers
}

resource "azuread_group" "cluster_viewer" {
  display_name            = "cluster-viewer"
  prevent_duplicate_names = true
  security_enabled        = true

  members = var.management_infra_cluster_viewer_owners
  owners  = var.management_infra_cluster_viewers
}
