locals {
  dns_zone          = coalesce(var.azure_foundation_public_zone_name, var.aws_foundation_public_zone_name, var.local_foundation_public_zone_name)
  oidc_url          = "oidc.${local.dns_zone}"
  oidc_callback_url = "oidc.${local.dns_zone}${var.local_foundation_public_zone_name != "" ? ":8443" : ""}/callback"

  foundation_name = coalesce(var.azure_foundation_name, var.aws_foundation_name, var.local_foundation_name)
}
