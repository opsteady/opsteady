locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_name": "${azurerm_kubernetes_cluster.platform.name}",
  "${local.component_name_underscores}_nodes_resource_group_name": "${azurerm_kubernetes_cluster.platform.node_resource_group}",
  "${local.component_name_underscores}_kubelet_identity_object_id": "${azurerm_kubernetes_cluster.platform.kubelet_identity.0.object_id}"
}
EOT
}
