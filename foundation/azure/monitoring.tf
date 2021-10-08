resource "azurerm_log_analytics_workspace" "platform" {
  name                = var.foundation_azure_name
  location            = azurerm_resource_group.foundation.location
  resource_group_name = azurerm_resource_group.foundation.name
  sku                 = "PerGB2018"
  retention_in_days   = var.foundation_azure_log_analytics_workspace_retention
}
