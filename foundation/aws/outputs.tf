locals {
  component_name_underscores = replace(var.platform_component_name, "-", "_")
}

resource "vault_generic_secret" "outputs" {
  path = "config/${var.platform_version}/platform/${var.platform_environment_name}/${var.platform_component_name}-tf"

  data_json = <<EOT
{
  "${local.component_name_underscores}_vpc_id": "${aws_vpc.platform.id}",
  "${local.component_name_underscores}_eks_a_subnet_id": "${aws_subnet.eks_a.id}",
  "${local.component_name_underscores}_eks_b_subnet_id": "${aws_subnet.eks_b.id}",
  "${local.component_name_underscores}_eks_c_subnet_id": "${aws_subnet.eks_c.id}",
  "${local.component_name_underscores}_pods_a_subnet_id": "${aws_subnet.pods_a.id}",
  "${local.component_name_underscores}_pods_b_subnet_id": "${aws_subnet.pods_b.id}",
  "${local.component_name_underscores}_pods_c_subnet_id": "${aws_subnet.pods_c.id}",
  "${local.component_name_underscores}_pods_a_cidr_block": "${aws_subnet.pods_a.cidr_block}",
  "${local.component_name_underscores}_pods_b_cidr_block": "${aws_subnet.pods_b.cidr_block}",
  "${local.component_name_underscores}_pods_c_cidr_block": "${aws_subnet.pods_c.cidr_block}",
  "${local.component_name_underscores}_public_zone_name": "${aws_route53_zone.public.name}",
  "${local.component_name_underscores}_public_zone_id": "${aws_route53_zone.public.zone_id}"
}
EOT
}
