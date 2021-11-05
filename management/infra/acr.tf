# Management ACR doesn't have any special requirements and can stay very simple
resource "azurerm_container_registry" "management" {
  name                = var.management_infra_acr_name
  resource_group_name = azurerm_resource_group.management.name
  location            = azurerm_resource_group.management.location
  sku                 = "Basic" # https://docs.microsoft.com/en-us/azure/container-registry/container-registry-skus
}

# Enable monitoring on the ACR https://docs.microsoft.com/en-us/azure/container-registry/monitor-service
resource "azurerm_monitor_diagnostic_setting" "acr" {
  name                       = "acr"
  target_resource_id         = azurerm_container_registry.management.id
  log_analytics_workspace_id = azurerm_log_analytics_workspace.management.id

  log {
    category = "ContainerRegistryRepositoryEvents"
    enabled  = true

    retention_policy {
      days    = 7 # In case of an incident we need time
      enabled = true
    }
  }

  log {
    category = "ContainerRegistryLoginEvents"
    enabled  = true

    retention_policy {
      days    = 7 # In case of an incident we need time
      enabled = true
    }
  }

  metric {
    category = "AllMetrics"
    enabled  = true

    retention_policy {
      days    = 2 # No long analysis is needed
      enabled = true
    }
  }
}

module "acr_management" {
  source = "../../internal/modules/service-principal"
  name   = "management-acr"
}

resource "azurerm_role_assignment" "acr_management_pull" {
  role_definition_name             = "AcrPull"
  scope                            = azurerm_container_registry.management.id
  principal_id                     = module.acr_management.azuread_service_principal_object_id
  skip_service_principal_aad_check = true # skip the Azure Active Directory check which may fail due to replication lag
}

resource "azurerm_role_assignment" "acr_management_push" {
  role_definition_name             = "AcrPush"
  scope                            = azurerm_container_registry.management.id
  principal_id                     = module.acr_management.azuread_service_principal_object_id
  skip_service_principal_aad_check = true # skip the Azure Active Directory check which may fail due to replication lag
}
