resource "azurerm_role_assignment" "crypto_user" {
  scope                = data.azurerm_key_vault.management.id
  role_definition_name = "Key Vault Crypto User"
  principal_id         = module.vault_unseal.azuread_service_principal_object_id
}
