# Management ACR doesn't have any special requirements and can stay very simple
resource "azurerm_container_registry" "management" {
  name                = var.management_infra_acr_name
  resource_group_name = azurerm_resource_group.management.name
  location            = azurerm_resource_group.management.location
  sku                 = "Basic" # https://docs.microsoft.com/en-us/azure/container-registry/container-registry-skus
}


# Enabling monitoring on the ACR https://docs.microsoft.com/en-us/azure/container-registry/monitor-service
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


# acr-management user is created to be used from EKS (AWS) to pull containers from management ACR
resource "azuread_application" "acr_management" {
  display_name               = "acr-management"
  oauth2_allow_implicit_flow = true
}

resource "azuread_service_principal" "acr_management" {
  application_id               = azuread_application.acr_management.application_id
  app_role_assignment_required = false
}

resource "azuread_service_principal_password" "acr_management" {
  service_principal_id = azuread_service_principal.acr_management.id
  description          = "ACR management password"
  value                = var.management_infra_acr_management_password
  end_date_relative    = "8760h"
}

resource "azurerm_role_assignment" "acr_management" {
  role_definition_name             = "AcrPull"
  scope                            = azurerm_container_registry.management.id
  principal_id                     = azuread_service_principal.acr_management.object_id
  skip_service_principal_aad_check = true # skip the Azure Active Directory check which may fail due to replication lag
}
