# This storage account will serve to host the Vault CA certificate. This
# certificate is used by the CI/CD mechanism to validate the Vault server.
resource "azurerm_storage_account" "vault_ca" {
  name                      = var.management_vault_infra_storage_account_name
  resource_group_name       = data.azurerm_resource_group.management.name
  location                  = var.management_infra_location
  account_kind              = "StorageV2"
  account_tier              = "Standard"
  account_replication_type  = "LRS"
  enable_https_traffic_only = true
  min_tls_version           = "TLS1_2"

  # To serve the Vault CA
  allow_blob_public_access = true

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_storage_container" "vault_ca" {
  name                  = "vault-ca"
  storage_account_name  = azurerm_storage_account.vault_ca.name
  container_access_type = "blob"
}

resource "azurerm_storage_blob" "vault_ca" {
  name                   = "ca.pem"
  storage_account_name   = azurerm_storage_account.vault_ca.name
  storage_container_name = azurerm_storage_container.vault_ca.name
  type                   = "Block"
  source_content         = tls_self_signed_cert.ca.cert_pem
}
