resource "vault_generic_secret" "outputs" {
  path = var.platform_terraform_output_path

  data_json = <<EOT
{
  "${var.platform_vault_vars_name}_vpc_id": "${aws_vpc.platform.id}",
  "${var.platform_vault_vars_name}_eks_a_subnet_id": "${aws_subnet.eks_a.id}",
  "${var.platform_vault_vars_name}_eks_b_subnet_id": "${aws_subnet.eks_b.id}",
  "${var.platform_vault_vars_name}_eks_c_subnet_id": "${aws_subnet.eks_c.id}",
  "${var.platform_vault_vars_name}_pods_a_subnet_id": "${aws_subnet.pods_a.id}",
  "${var.platform_vault_vars_name}_pods_b_subnet_id": "${aws_subnet.pods_b.id}",
  "${var.platform_vault_vars_name}_pods_c_subnet_id": "${aws_subnet.pods_c.id}",
  "${var.platform_vault_vars_name}_pods_a_cidr_block": "${aws_subnet.pods_a.cidr_block}",
  "${var.platform_vault_vars_name}_pods_b_cidr_block": "${aws_subnet.pods_b.cidr_block}",
  "${var.platform_vault_vars_name}_pods_c_cidr_block": "${aws_subnet.pods_c.cidr_block}",
  "${var.platform_vault_vars_name}_public_zone_name": "${aws_route53_zone.public.name}",
  "${var.platform_vault_vars_name}_public_zone_id": "${aws_route53_zone.public.zone_id}"
}
EOT
}
