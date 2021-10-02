resource "azuread_application" "spn" {
  display_name            = var.name
  owners                  = var.owners
  group_membership_claims = var.group_membership_claims

  dynamic "required_resource_access" {
    for_each = var.required_resource_access
    content {
      resource_app_id = required_resource_access.value["resource_app_id"]

      dynamic "resource_access" {
        for_each = required_resource_access.value["resource_access"]
        content {
          id   = resource_access.value["id"]
          type = resource_access.value["type"]
        }
      }
    }
  }

  dynamic "app_role" {
    for_each = var.app_roles
    content {
      id                   = app_role.value["id"]
      allowed_member_types = app_role.value["allowed_member_types"]
      description          = app_role.value["description"]
      display_name         = app_role.value["display_name"]
      enabled              = app_role.value["enabled"]
      value                = app_role.value["value"]
    }
  }

  web {
    redirect_uris = var.redirect_uris
  }
}

resource "azuread_service_principal" "spn" {
  application_id               = azuread_application.spn.application_id
  app_role_assignment_required = false
}

resource "random_password" "spn" {
  length  = 24
  special = true
}

resource "time_rotating" "spn" {
  rotation_days = 30
}

resource "azuread_service_principal_password" "spn" {
  service_principal_id = azuread_service_principal.spn.id

  rotate_when_changed = {
    rotation = time_rotating.spn.id
  }
}
