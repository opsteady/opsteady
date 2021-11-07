locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
  azure_dns_config = <<EOF
{
  "tenantId": "${data.azurerm_client_config.current.tenant_id}",
  "subscriptionId": "${var.foundation_azure_subscription_id}",
  "resourceGroup": "${azurerm_resource_group.dns.name}",
  "useManagedIdentityExtension": true
}
EOF
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_msi_id": "${azurerm_user_assigned_identity.dns.id}",
  "${local.component_name_underscores}_msi_client_id": "${azurerm_user_assigned_identity.dns.client_id}",
  "${local.component_name_underscores}_azure_dns_config": "${base64encode(local.azure_dns_config)}"
}
EOT
}
