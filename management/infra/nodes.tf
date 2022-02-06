resource "azurerm_kubernetes_cluster_node_pool" "platform" {
  count = var.management_infra_minimal ? 0 : 1

  name                   = "platform"
  kubernetes_cluster_id  = azurerm_kubernetes_cluster.management.0.id
  vm_size                = var.management_infra_aks_platform_nodepool_node_size
  node_count             = var.management_infra_aks_platform_nodepool_node_count
  vnet_subnet_id         = azurerm_subnet.pods.id
  enable_host_encryption = var.management_infra_host_encryption
  mode                   = "System"
  availability_zones     = ["1", "2", "3"]
}
