resource "azurerm_role_assignment" "aks_msi_managed_identity_operator" {
  role_definition_name = "Managed Identity Operator"
  principal_id         = var.azure_cluster_kubelet_identity_object_id
  scope                = "/subscriptions/${var.azure_foundation_subscription_id}/resourcegroups/${var.azure_cluster_nodes_resource_group_name}"
}

resource "azurerm_role_assignment" "aks_msi_virtual_machine_contributor" {
  role_definition_name = "Virtual Machine Contributor"
  principal_id         = var.azure_cluster_kubelet_identity_object_id
  scope                = "/subscriptions/${var.azure_foundation_subscription_id}/resourcegroups/${var.azure_cluster_nodes_resource_group_name}"
}
