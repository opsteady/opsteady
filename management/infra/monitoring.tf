# Used to store logs from various Azure services for debugging and insights
resource "azurerm_log_analytics_workspace" "management" {
  name                = "management-analytics"
  location            = azurerm_resource_group.management.location
  resource_group_name = azurerm_resource_group.management.name
  sku                 = "PerGB2018" # Using the default
  retention_in_days   = var.management_infra_log_analytics_workspace_retention
}
