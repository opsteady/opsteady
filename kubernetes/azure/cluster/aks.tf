resource "azurerm_kubernetes_cluster" "platform" {
  name                      = var.foundation_azure_name
  dns_prefix                = var.foundation_azure_name
  node_resource_group       = "platform-nodes-${var.foundation_azure_name}"
  location                  = azurerm_resource_group.kubernetes.location
  resource_group_name       = azurerm_resource_group.kubernetes.name
  sku_tier                  = "Free"
  kubernetes_version        = var.kubernetes_azure_cluster_kubernetes_version
  private_cluster_enabled   = false
  disk_encryption_set_id    = azurerm_disk_encryption_set.platform.id

  default_node_pool {
    name                   = "system"
    vnet_subnet_id         = var.foundation_azure_pods_subnet_id
    node_count             = var.kubernetes_azure_cluster_system_nodepool_node_count
    vm_size                = var.kubernetes_azure_cluster_system_nodepool_node_size
    orchestrator_version   = var.kubernetes_azure_cluster_kubernetes_version
    enable_auto_scaling    = false
    availability_zones     = ["1", "2", "3"]
    enable_host_encryption = true
  }

  network_profile {
    network_plugin     = "azure"
    network_policy     = "azure"
    service_cidr       = "172.20.0.0/16"
    docker_bridge_cidr = "172.21.0.1/16"
    dns_service_ip     = "172.20.0.10"
    load_balancer_sku  = "Standard"
    outbound_type      = "loadBalancer"
    load_balancer_profile {
      outbound_ip_prefix_ids = [var.foundation_azure_public_ip_prefix_id]
    }
  }

  identity {
    type = "SystemAssigned"
  }

  addon_profile {
    kube_dashboard {
      enabled = false
    }
  }

  role_based_access_control {
    enabled = true
  }
}
