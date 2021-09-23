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
    namespace = "platform"
  }

  data = {
    "tls.crt" = tls_locally_signed_cert.vault.cert_pem,
    "tls.key" = tls_private_key.vault.private_key_pem,
    "ca.crt"  = tls_self_signed_cert.ca.cert_pem
  }
}
