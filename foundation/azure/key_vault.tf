# Used for storing certificates or CMKs when needed by the platform
# Not used in the cluster for storing secrets
resource "azurerm_key_vault" "platform" {
  name                        = var.foundation_azure_name
  location                    = azurerm_resource_group.foundation.location
  resource_group_name         = azurerm_resource_group.foundation.name
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  purge_protection_enabled    = true
  enable_rbac_authorization   = true
  enabled_for_disk_encryption = true
  sku_name                    = "standard"
}
