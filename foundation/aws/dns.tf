# Create a unique subdomain for this platform
resource "aws_route53_zone" "public" {
  name = "${var.aws_foundation_public_name}.${var.management_infra_domain}"
}

resource "azurerm_dns_ns_record" "public_management_delegation" {
  provider = azurerm.management

  name                = var.aws_foundation_public_name
  zone_name           = data.azurerm_dns_zone.public_root.name
  resource_group_name = "management"
  ttl                 = 300

  records = aws_route53_zone.public.name_servers
}
