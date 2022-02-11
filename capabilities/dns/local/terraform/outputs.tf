locals {
  azure_dns_config = jsonencode({
    "tenantId" : "${data.azurerm_client_config.current.tenant_id}",
    "subscriptionId" : "${var.local_foundation_subscription_id}",
    "resourceGroup" : "${azurerm_resource_group.dns.name}",
    "aadClientId" : "${module.dns.azuread_service_principal_application_id}",
    "aadClientSecret" : "${module.dns.azuread_service_principal_password}"
  })
}

resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${local.component_name_underscores}_azure_dns_config": "${base64encode(local.azure_dns_config)}",
  "${local.component_name_underscores}_resource_group_name": "${azurerm_resource_group.dns.name}"
}
EOT
}
