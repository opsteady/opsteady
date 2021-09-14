module "vault_unseal" {
  source = "../../internal/modules/service-principal"
  name   = "vault-unseal"
}

resource "tls_private_key" "vault" {
  algorithm = "RSA"
  rsa_bits  = "2048"
}

resource "tls_cert_request" "vault" {
  key_algorithm   = "RSA"
  private_key_pem = tls_private_key.vault.private_key_pem

  subject {
    common_name = "vault"
  }

  dns_names = [
    "*.vault-internal",
    "vault.management.${var.management_infra_domain}"
  ]

  ip_addresses = [
    "127.0.0.1"
  ]
}

resource "tls_locally_signed_cert" "vault" {
  cert_request_pem   = tls_cert_request.vault.cert_request_pem
  ca_key_algorithm   = "RSA"
  ca_private_key_pem = tls_private_key.ca.private_key_pem
  ca_cert_pem        = tls_self_signed_cert.ca.cert_pem

  validity_period_hours = 8640

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "kubernetes_secret" "vault_tls" {
  metadata {
    name      = "vault-tls"
    namespace = kubernetes_namespace.platform.metadata.0.name
  }

  data = {
    "tls.crt" = tls_locally_signed_cert.vault.cert_pem,
    "tls.key" = tls_private_key.vault.private_key_pem,
    "ca.crt"  = tls_self_signed_cert.ca.cert_pem
  }
}

resource "helm_release" "vault" {
  name       = "vault"
  repository = "https://helm.releases.hashicorp.com/"
  chart      = "vault"
  namespace  = kubernetes_namespace.platform.metadata.0.name
  version    = var.management_infra_vault_chart_version

  values = [
    templatefile("templates/values.yaml.tmpl", {
      replicas               = 3,
      vault_affinity         = "" # TODO: remove for true HA
      vault_image_repository = var.management_infra_vault_image_repository,
      vault_image_tag        = var.management_infra_vault_image_tag,
      client_id              = module.vault_unseal.azuread_application_application_id,
      client_secret          = module.vault_unseal.azuread_service_principal_password
      vault_name             = azurerm_key_vault.management.name
      key_name               = azurerm_key_vault_key.vault.name
      tenant_id              = data.azurerm_client_config.current.tenant_id
      domain                 = var.management_infra_domain
    })
  ]

  depends_on = [kubernetes_secret.vault_tls]
}
