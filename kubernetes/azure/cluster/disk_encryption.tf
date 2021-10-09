resource "azurerm_key_vault_key" "platform" {
  name         = "disk-enc-${var.foundation_azure_name}"
  key_vault_id = var.foundation_azure_key_vault_id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "unwrapKey",
    "wrapKey"
  ]
}

resource "azurerm_disk_encryption_set" "platform" {
  name                = var.foundation_azure_name
  resource_group_name = azurerm_resource_group.kubernetes.name
  location            = azurerm_resource_group.kubernetes.location
  key_vault_key_id    = azurerm_key_vault_key.platform.id

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_role_assignment" "disk_encryption_crypto_user" {
  scope                = var.foundation_azure_key_vault_id
  role_definition_name = "Key Vault Crypto Service Encryption User"
  principal_id         = azurerm_disk_encryption_set.platform.identity.0.principal_id
}

resource "azurerm_role_assignment" "aks_cluster_disk_encryption_reader" {
  scope                = azurerm_disk_encryption_set.platform.id
  role_definition_name = "Reader"
  principal_id         = azurerm_kubernetes_cluster.platform.identity.0.principal_id
}
