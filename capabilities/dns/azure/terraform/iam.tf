# Managed User Identity for external-dns
resource "azurerm_user_assigned_identity" "dns" {
  resource_group_name = azurerm_resource_group.dns.name
  location            = azurerm_resource_group.dns.location

  name = "dns"
}

# Allow MSI to read from the DNS resource group
resource "azurerm_role_assignment" "dns_resource_group_reader" {
  role_definition_name = "Reader"
  principal_id         = azurerm_user_assigned_identity.dns.principal_id
  scope                = "/subscriptions/${data.azurerm_client_config.current.subscription_id}/resourcegroups/${var.foundation_azure_resource_group}"
}

# Allow MSI to modify public DNS zone
resource "azurerm_role_assignment" "dns_zone_contributor" {
  role_definition_name = "Contributor"
  principal_id         = azurerm_user_assigned_identity.dns.principal_id
  scope                = var.foundation_azure_public_zone_id
}

# Allow the Kubelet identity to use/assign DNS MSI
resource "azurerm_role_assignment" "aks_msi_managed_identity_operator_dns_msi" {
  role_definition_name = "Managed Identity Operator"
  principal_id         = var.kubernetes_azure_cluster_kubelet_identity_object_id
  scope                = azurerm_user_assigned_identity.dns.id
}
