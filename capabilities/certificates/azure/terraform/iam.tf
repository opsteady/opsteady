# Managed User Identity for external-dns
resource "azurerm_user_assigned_identity" "certificates" {
  resource_group_name = azurerm_resource_group.certificates.name
  location            = azurerm_resource_group.certificates.location

  name = "certificates"
}


# Allow MSI to modify public DNS zone
resource "azurerm_role_assignment" "dns_zone_contributor" {
  role_definition_name = "Contributor"
  principal_id         = azurerm_user_assigned_identity.certificates.principal_id
  scope                = var.foundation_azure_public_zone_id
}

# Allow the Kubelet identity to use/assign certificates MSI
resource "azurerm_role_assignment" "aks_msi_managed_identity_operator_certificates_msi" {
  role_definition_name = "Managed Identity Operator"
  principal_id         = var.kubernetes_azure_cluster_kubelet_identity_object_id
  scope                = azurerm_user_assigned_identity.certificates.id
}
