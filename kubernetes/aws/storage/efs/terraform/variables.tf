variable "foundation_aws_name" {
  type = string
}

variable "foundation_aws_region" {
  type = string
}

variable "foundation_aws_vpc_id" {
  type = string
}

variable "foundation_aws_pods_a_subnet_id" {
  type = string
}

variable "foundation_aws_pods_b_subnet_id" {
  type = string
}

variable "foundation_aws_pods_c_subnet_id" {
  type = string
}

variable "foundation_aws_pods_a_cidr_block" {
  type = string
}

variable "foundation_aws_pods_b_cidr_block" {
  type = string
}

variable "foundation_aws_pods_c_cidr_block" {
  type = string
}

variable "kubernetes_aws_cluster_security_group_id" {
  type = string
}

variable "kubernetes_aws_cluster_openid_connect_provider_platform_arn" {
  type = string
}

variable "kubernetes_aws_cluster_openid_connect_provider_platform_url" {
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

variable "vault_address" {
  type = string
}

variable "vault_token" {
  type = string
}
