locals {
  dns_zone          = coalesce(var.foundation_azure_public_zone_name, var.foundation_aws_public_zone_name, var.foundation_local_public_zone_name)
  oidc_url          = "oidc.${local.dns_zone}"
  oidc_callback_url = "oidc.${local.dns_zone}${var.foundation_local_public_zone_name != "" ? ":8443" : ""}/callback"

  foundation_name = coalesce(var.foundation_azure_name, var.foundation_aws_name, var.foundation_local_name)
}
