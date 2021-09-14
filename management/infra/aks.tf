resource "azurerm_kubernetes_cluster" "management" {
  name                            = var.management_infra_aks_name
  dns_prefix                      = var.management_infra_aks_name
  node_resource_group             = "nodes-${var.management_infra_aks_name}"
  location                        = azurerm_resource_group.management.location
  resource_group_name             = azurerm_resource_group.management.name
  sku_tier                        = var.management_infra_aks_sku_tier
  kubernetes_version              = var.management_infra_aks_kubernetes_version
  api_server_authorized_ip_ranges = var.management_infra_aks_api_server_authorized_ip_ranges
  disk_encryption_set_id          = azurerm_disk_encryption_set.management.id
  private_cluster_enabled         = false

  default_node_pool {
    name                 = "system"
    vnet_subnet_id       = azurerm_subnet.pods.id
    node_count           = var.management_infra_aks_system_nodepool_node_count
    vm_size              = var.management_infra_aks_system_nodepool_node_size
    orchestrator_version = var.management_infra_aks_kubernetes_version
    enable_auto_scaling  = false
    availability_zones   = ["1", "2", "3"]
  }

  network_profile {
    network_plugin     = "azure"
    network_policy     = "calico"
    service_cidr       = "172.20.0.0/16"
    docker_bridge_cidr = "172.21.0.1/16"
    dns_service_ip     = "172.20.0.10"
    load_balancer_sku  = "Standard"
    outbound_type      = "loadBalancer"
    load_balancer_profile {
      outbound_ip_prefix_ids = [azurerm_public_ip_prefix.pub.id]
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

resource "azurerm_monitor_diagnostic_setting" "aks" {
  name                       = var.management_infra_aks_name
  target_resource_id         = azurerm_kubernetes_cluster.management.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.management.id

  log {
    category = "kube-apiserver"
    enabled  = true

    retention_policy {
      enabled = false
    }
  }

  log {
    category = "kube-controller-manager"
    enabled  = true

    retention_policy {
      enabled = false
    }
  }
}
