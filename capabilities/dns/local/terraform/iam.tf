# Service principal for external-dns
module "dns" {
  source = "../../../../internal/modules/service-principal"
  name   = "dns-${var.foundation_local_name}"
}

# Allow SP to read from the DNS resource group
resource "azurerm_role_assignment" "dns_resource_group_reader" {
  role_definition_name = "Reader"
  principal_id         = module.dns.azuread_service_principal_object_id
  scope                = "/subscriptions/${data.azurerm_client_config.current.subscription_id}/resourcegroups/${var.foundation_local_resource_group}"
}

# Allow SP to modify public DNS zone
resource "azurerm_role_assignment" "dns_zone_contributor" {
  role_definition_name = "Contributor"
  principal_id         = module.dns.azuread_service_principal_object_id
  scope                = var.foundation_local_public_zone_id
}
