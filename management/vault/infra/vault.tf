module "vault_unseal" {
  source = "../../../internal/modules/service-principal"
  name   = "vault-unseal-new"
}

# Vault auto-unseal key

resource "azurerm_key_vault_key" "vault" {
  name         = "vault"
  key_vault_id = data.azurerm_key_vault.management.id
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "wrapKey",
    "unwrapKey",
  ]
}

# We are using Terraform to deploy Helm via the Helm chart. The primary reason for doing this
# is the tight connection between Terraform resource attributes and Vault configuration.
# Furthermore, we expect relatively few changes to the configuration.
resource "helm_release" "vault" {
  name       = "vault"
  repository = "https://helm.releases.hashicorp.com/"
  chart      = "vault"
  namespace  = "platform"
  version    = var.management_vault_infra_chart_version

  # The Vault pods will be in a sealed state on first start, so don't wait for them
  # to become ready.
  wait = false

  values = [
    templatefile("templates/values.yaml.tmpl", {
      vault_image_repository = var.management_vault_infra_image_repository,
      vault_image_tag        = var.management_vault_infra_image_tag,
      client_id              = module.vault_unseal.azuread_application_application_id,
      client_secret          = module.vault_unseal.azuread_service_principal_password
      vault_name             = data.azurerm_key_vault.management.name
      key_name               = azurerm_key_vault_key.vault.name
      tenant_id              = data.azurerm_client_config.current.tenant_id
      domain                 = var.management_infra_domain
      loadbalancer_ip        = azurerm_public_ip.vault.ip_address
    })
  ]

  depends_on = [kubernetes_secret.vault_tls]
}
