// Used only for storing terraform state
// Azure Storage encrypts data when it is persisted to the cloud, we don't think a private key is needed
resource "azurerm_storage_account" "terraform_state" {
  name                      = var.management_bootstrap_terraform_state_account_name
  resource_group_name       = azurerm_resource_group.terraform_state.name
  location                  = azurerm_resource_group.terraform_state.location
  account_kind              = "StorageV2"
  account_tier              = "Standard" // Used for tf storage, requirement is low speed and low volume
  account_replication_type  = "LRS"      // 11 nines
  enable_https_traffic_only = true
  blob_properties {
    versioning_enabled = true
    delete_retention_policy {
      days = 365
    }
  }
}

# Used for management terraform state
resource "azurerm_storage_container" "management" {
  name                 = "management"
  storage_account_name = azurerm_storage_account.terraform_state.name
}

# Used for all user platform terraform state
resource "azurerm_storage_container" "platform" {
  name                 = "platform"
  storage_account_name = azurerm_storage_account.terraform_state.name
}
