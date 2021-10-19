resource "azurerm_monitor_diagnostic_setting" "k8s" {
  name                       = "k8s"
  target_resource_id         = azurerm_kubernetes_cluster.platform.id
  log_analytics_workspace_id = var.foundation_azure_log_analytics_id

  log {
    category = "kube-scheduler"
    enabled  = false

    retention_policy {
      enabled = false
    }
  }

  log {
    category = "kube-controller-manager"
    enabled  = true

    retention_policy {
      enabled = true
      days    = 5
    }
  }

  log {
    category = "cluster-autoscaler"
    enabled  = false

    retention_policy {
      enabled = false
    }
  }

  log {
    category = "kube-audit"
    enabled  = true

    retention_policy {
      enabled = true
      days    = 5
    }
  }

  log {
    category = "kube-apiserver"
    enabled  = true

    retention_policy {
      enabled = true
      days    = 5
    }
  }
}
