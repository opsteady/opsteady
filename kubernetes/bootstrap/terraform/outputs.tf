locals {
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
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_platform_admin_group_object_id": "${data.azuread_group.platform_admin.object_id}",
  "${var.platform_vault_vars_name}_platform_operator_group_object_id": "${data.azuread_group.platform_operator.object_id}",
  "${var.platform_vault_vars_name}_platform_viewer_group_object_id": "${data.azuread_group.platform_viewer.object_id}",
  "${var.platform_vault_vars_name}_management_acr_docker_config": "${base64encode(local.management_acr_docker_config)}"
}
EOT
}
