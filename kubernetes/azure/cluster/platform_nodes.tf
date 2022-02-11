resource "azurerm_kubernetes_cluster_node_pool" "platform" {
  name                   = "platform"
  kubernetes_cluster_id  = azurerm_kubernetes_cluster.platform.id
  vm_size                = var.azure_cluster_platform_nodepool_node_size
  node_count             = var.azure_cluster_platform_nodepool_node_count
  vnet_subnet_id         = var.foundation_azure_pods_subnet_id
  enable_host_encryption = true
  mode                   = "System"
  availability_zones     = ["1", "2", "3"]
  node_labels            = { name : "system" }
}
