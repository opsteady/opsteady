resource "azurerm_key_vault" "management" {
  name                      = var.management_infra_key_vault_name
  location                  = azurerm_resource_group.management.location
  resource_group_name       = azurerm_resource_group.management.name
  tenant_id                 = data.azurerm_client_config.current.tenant_id
  purge_protection_enabled  = true
  enable_rbac_authorization = true
  sku_name                  = "standard"

  network_acls {
    default_action = "Deny"
    bypass         = "AzureServices"

    ip_rules = var.management_infra_key_vault_ip_rules
    virtual_network_subnet_ids = [azurerm_subnet.pods.id] # Needed to give the Vault pods access to Key Vault
  }
}

resource "azurerm_role_assignment" "key_vault_administrator" {
  for_each = toset(var.management_infra_key_vault_administrators)

  scope                = azurerm_key_vault.management.id
  role_definition_name = "Key Vault Administrator"
  principal_id         = each.value
}

resource "time_sleep" "wait_for_key_vault_iam_propagation" {
  depends_on = [azurerm_role_assignment.key_vault_administrator]

  create_duration = "10s"
}

# Disk encryption resources for AKS

resource "azurerm_key_vault_key" "management" {
  name         = var.management_infra_aks_name
  key_vault_id = azurerm_key_vault.management.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "unwrapKey",
    "wrapKey"
  ]

  depends_on = [time_sleep.wait_for_key_vault_iam_propagation]
}

resource "azurerm_disk_encryption_set" "management" {
  name                = var.management_infra_aks_name
  resource_group_name = azurerm_resource_group.management.name
  location            = azurerm_resource_group.management.location
  key_vault_key_id    = azurerm_key_vault_key.management.id

  identity {
    type = "SystemAssigned"
  }
}

resource "azurerm_role_assignment" "disk_encryption_crypto_user" {
  scope                = azurerm_key_vault.management.id
  role_definition_name = "Key Vault Crypto Service Encryption User"
  principal_id         = azurerm_disk_encryption_set.management.identity.0.principal_id
}

resource "azurerm_role_assignment" "aks_cluster_disk_encryption_reader" {
  scope                = azurerm_disk_encryption_set.management.id
  role_definition_name = "Reader"
  principal_id         = azurerm_kubernetes_cluster.management.identity.0.principal_id
}
