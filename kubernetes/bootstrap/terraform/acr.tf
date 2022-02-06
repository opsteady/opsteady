data "azuread_application" "management_acr" {
  display_name = "management-acr"
}

resource "time_rotating" "management_acr" {
  rotation_days = 7
}

resource "azuread_application_password" "management_acr" {
  application_object_id = data.azuread_application.management_acr.object_id
  display_name          = coalesce(var.foundation_aws_environment_name, var.foundation_azure_environment_name, var.foundation_local_environment_name)
  rotate_when_changed = {
    rotation = time_rotating.management_acr.id
  }
}
