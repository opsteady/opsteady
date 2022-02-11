resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_name": "${azurerm_kubernetes_cluster.platform.name}",
  "${var.platform_vault_vars_name}_nodes_resource_group_name": "${azurerm_kubernetes_cluster.platform.node_resource_group}",
  "${var.platform_vault_vars_name}_kubelet_identity_object_id": "${azurerm_kubernetes_cluster.platform.kubelet_identity.0.object_id}"
}
EOT
}
