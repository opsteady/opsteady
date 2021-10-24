locals {
  component_name_underscores   = replace(var.platform_component_name, "-", "_")
  management_acr_docker_config = <<DOCKER
{
  "auths": {
    "${var.management_infra_acr_name}.azurecr.io": {
      "auth": "${base64encode("${data.azuread_application.management_acr.application_id}:${azuread_application_password.management_acr.value}")}"
    }
  }
}
DOCKER
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_platform_admin_group_object_id": "${data.azuread_group.platform_admin.object_id}",
  "${local.component_name_underscores}_platform_operator_group_object_id": "${data.azuread_group.platform_operator.object_id}",
  "${local.component_name_underscores}_platform_viewer_group_object_id": "${data.azuread_group.platform_viewer.object_id}",
  "${local.component_name_underscores}_management_acr_docker_config": "${base64encode(local.management_acr_docker_config)}"
}
EOT
}
