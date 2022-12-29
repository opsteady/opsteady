variable "aws_foundation_name" {
  description = "Name to be used for resources or as a suffix, mostly plt1"
  type        = string
}

variable "aws_foundation_environment_name" {
  description = "Name of the platform environment, for example dev-azure"
  type        = string
}

variable "aws_foundation_public_name" {
  description = "The name used as the sub domain"
  type        = string
}

variable "aws_foundation_region" {
  type = string
}

variable "aws_foundation_vpc_cidr" {
  type = string
}

variable "aws_foundation_subnet_eks_a" {
  type = string
}

variable "aws_foundation_subnet_eks_b" {
  type = string
}

variable "aws_foundation_subnet_eks_c" {
  type = string
}

variable "aws_foundation_zone_eks_a" {
  type = string
}

variable "aws_foundation_zone_eks_b" {
  type = string
}

variable "aws_foundation_zone_eks_c" {
  type = string
}

variable "aws_foundation_subnet_pods_a" {
  type = string
}

variable "aws_foundation_subnet_pods_b" {
  type = string
}

variable "aws_foundation_subnet_pods_c" {
  type = string
}

variable "aws_foundation_zone_pods_a" {
  type = string
}

variable "aws_foundation_zone_pods_b" {
  type = string
}

variable "aws_foundation_zone_pods_c" {
  type = string
}

variable "aws_foundation_subnet_pub_a" {
  type = string
}

variable "aws_foundation_subnet_pub_b" {
  type = string
}

variable "aws_foundation_subnet_pub_c" {
  type = string
}

variable "aws_foundation_subnet_prv_a" {
  type = string
}

variable "aws_foundation_subnet_prv_b" {
  type = string
}

variable "aws_foundation_subnet_prv_c" {
  type = string
}

variable "aws_foundation_zone_pub_a" {
  type = string
}

variable "aws_foundation_zone_pub_b" {
  type = string
}

variable "aws_foundation_zone_pub_c" {
  type = string
}

variable "aws_foundation_zone_prv_a" {
  type = string
}

variable "aws_foundation_zone_prv_b" {
  type = string
}

variable "aws_foundation_zone_prv_c" {
  type = string
}

variable "aws_foundation_nat_a_enabeld" {
  type = bool
}

variable "aws_foundation_nat_b_enabeld" {
  type = bool
}

variable "aws_foundation_nat_c_enabeld" {
  type = bool
}

variable "management_infra_domain" {
  type = string
}

variable "management_subscription_id" {
  type = string
}

variable "management_client_id" {
  type = string
}

variable "management_client_secret" {
  type = string
}

variable "tenant_id" {
  type = string
}

variable "platform_terraform_output_path" {
  type = string
}

variable "platform_vault_vars_name" {
  type = string
}

variable "vault_address" {
  type = string
}

variable "vault_token" {
  type = string
}
