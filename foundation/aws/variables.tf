variable "foundation_aws_name" {
  description = "Name to be used for resources or as a suffix, mostly plt1"
  type = string
}

variable "foundation_aws_environment_name" {
  description = "Name of the platform environment, for example dev-azure"
  type = string
}

variable "foundation_aws_public_name" {
  description = "The name used as the sub domain"
  type = string
}

variable "foundation_aws_region" {
  type = string
}

variable "foundation_aws_vpc_cidr" {
  type = string
}

variable "foundation_aws_subnet_eks_a" {
  type = string
}

variable "foundation_aws_subnet_eks_b" {
  type = string
}

variable "foundation_aws_subnet_eks_c" {
  type = string
}

variable "foundation_aws_zone_eks_a" {
  type = string
}

variable "foundation_aws_zone_eks_b" {
  type = string
}

variable "foundation_aws_zone_eks_c" {
  type = string
}

variable "foundation_aws_subnet_pods_a" {
  type = string
}

variable "foundation_aws_subnet_pods_b" {
  type = string
}

variable "foundation_aws_subnet_pods_c" {
  type = string
}

variable "foundation_aws_zone_pods_a" {
  type = string
}

variable "foundation_aws_zone_pods_b" {
  type = string
}

variable "foundation_aws_zone_pods_c" {
  type = string
}

variable "foundation_aws_subnet_pub_a" {
  type = string
}

variable "foundation_aws_subnet_pub_b" {
  type = string
}

variable "foundation_aws_subnet_pub_c" {
  type = string
}

variable "foundation_aws_zone_pub_a" {
  type = string
}

variable "foundation_aws_zone_pub_b" {
  type = string
}

variable "foundation_aws_zone_pub_c" {
  type = string
}

variable "foundation_aws_nat_a_enabeld" {
  type = bool
}

variable "foundation_aws_nat_b_enabeld" {
  type = bool
}

variable "foundation_aws_nat_c_enabeld" {
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

variable "platform_version" {
  type = string
}

variable "platform_environment_name" {
  type = string
}

variable "platform_component_name" {
  type = string
}

# Used for creating output to Vault
variable "vault_address" {
  type = string
}

variable "vault_token" {
  type = string
}
