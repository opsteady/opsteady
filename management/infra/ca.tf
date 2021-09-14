resource "tls_private_key" "ca" {
  algorithm = "RSA"
  rsa_bits  = "2048"
}

resource "tls_self_signed_cert" "ca" {
  key_algorithm   = tls_private_key.ca.algorithm
  private_key_pem = tls_private_key.ca.private_key_pem

  validity_period_hours = 8640

  is_ca_certificate = true

  allowed_uses = [
    "cert_signing",
    "key_encipherment",
    "digital_signature",
  ]

  subject {
    common_name = "Platform CA"
  }
}
