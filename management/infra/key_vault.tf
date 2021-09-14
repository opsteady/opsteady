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
  }
}

resource "azurerm_role_assignment" "crypto_user" {
  scope                = azurerm_key_vault.management.id
  role_definition_name = "Key Vault Crypto User"
  principal_id         = module.vault_unseal.azuread_service_principal_object_id
}

resource "azurerm_private_endpoint" "keyvault_pods" {
  name                = "keyvault-pods"
  resource_group_name = azurerm_resource_group.management.name
  location            = azurerm_resource_group.management.location
  subnet_id           = azurerm_subnet.pods.id

  private_service_connection {
    name                           = "vault-pods"
    private_connection_resource_id = azurerm_key_vault.management.id
    subresource_names              = ["Vault"]
    is_manual_connection           = false
  }
}

resource "azurerm_private_dns_zone" "keyvault" {
  name                = "privatelink.vaultcore.azure.net"
  resource_group_name = azurerm_resource_group.management.name
}

resource "azurerm_private_dns_a_record" "keyvault" {
  name                = azurerm_key_vault.management.name
  zone_name           = azurerm_private_dns_zone.keyvault.name
  resource_group_name = azurerm_resource_group.management.name
  ttl                 = 300
  records = [
    azurerm_private_endpoint.keyvault_pods.private_service_connection.0.private_ip_address
  ]
}

resource "azurerm_private_dns_zone_virtual_network_link" "keyvault" {
  name                  = "keyvault"
  resource_group_name   = azurerm_resource_group.management.name
  private_dns_zone_name = azurerm_private_dns_zone.keyvault.name
  virtual_network_id    = azurerm_virtual_network.management.id
}

resource "azurerm_role_assignment" "key_vault_administrator" {
  for_each = toset(concat([
    data.azurerm_client_config.current.object_id],
    var.management_infra_key_vault_administrators
  ))

  scope                = azurerm_key_vault.management.id
  role_definition_name = "Key Vault Administrator"
  principal_id         = each.value
}

resource "time_sleep" "wait_for_key_vault_iam_propagation" {
  depends_on = [azurerm_role_assignment.key_vault_administrator]

  create_duration = "10s"
}

# Vault auto-unseal key

resource "azurerm_key_vault_key" "vault" {
  name         = "vault"
  key_vault_id = azurerm_key_vault.management.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "wrapKey",
    "unwrapKey",
  ]

  depends_on = [time_sleep.wait_for_key_vault_iam_propagation]
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
