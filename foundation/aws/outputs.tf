locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

# resource "vault_generic_secret" "outputs" {
#   path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

#   data_json = <<EOT
# {
#   // TODO: what to upload?
# }
# EOT
# }
