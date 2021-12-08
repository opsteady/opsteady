locals {
  dns_zone          = var.foundation_azure_public_zone_name != "" ? var.foundation_azure_public_zone_name : var.foundation_aws_public_zone_name
  oidc_url          = "oidc.${local.dns_zone}"
  oidc_callback_url = "oidc.${local.dns_zone}/callback"

  foundation_name = var.foundation_azure_name != "" ? var.foundation_azure_name : var.foundation_aws_name
}
