# Service principal for certificate manager
module "certificates" {
  source = "../../../../internal/modules/service-principal"
  name   = "certificates-${var.foundation_local_name}"
}

# Allow SP to modify public DNS zone
resource "azurerm_role_assignment" "dns_zone_contributor" {
  role_definition_name = "Contributor"
  principal_id         = module.certificates.azuread_service_principal_object_id
  scope                = var.foundation_local_public_zone_id
}
