locals {
  azure_dns_config = <<EOF
{
  "tenantId": "${data.azurerm_client_config.current.tenant_id}",
  "subscriptionId": "${var.azure_foundation_subscription_id}",
  "resourceGroup": "${azurerm_resource_group.dns.name}",
  "useManagedIdentityExtension": true
}
EOF
}

resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_msi_id": "${azurerm_user_assigned_identity.dns.id}",
  "${var.platform_vault_vars_name}_msi_client_id": "${azurerm_user_assigned_identity.dns.client_id}",
  "${var.platform_vault_vars_name}_azure_dns_config": "${base64encode(local.azure_dns_config)}",
  "${var.platform_vault_vars_name}_resource_group_name": "${azurerm_resource_group.dns.name}"
}
EOT
}
