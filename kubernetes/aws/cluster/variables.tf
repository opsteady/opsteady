variable "aws_foundation_region" {
  type = string
}

variable "aws_foundation_name" {
  type = string
}

variable "aws_foundation_vpc_id" {
  type = string
}

variable "aws_foundation_eks_a_subnet_id" {
  type = string
}

variable "aws_foundation_eks_b_subnet_id" {
  type = string
}

variable "aws_foundation_eks_c_subnet_id" {
  type = string
}

variable "aws_foundation_pods_a_subnet_id" {
  type = string
}

variable "aws_foundation_pods_b_subnet_id" {
  type = string
}

variable "aws_foundation_pods_c_subnet_id" {
  type = string
}

variable "aws_foundation_pods_a_cidr_block" {
  type = string
}

variable "aws_foundation_pods_b_cidr_block" {
  type = string
}

variable "aws_foundation_pods_c_cidr_block" {
  type = string
}

variable "aws_cluster_public_access_cidrs" {
  type = list(string)
}

variable "aws_cluster_kubernetes_version" {
  type = string
}

variable "aws_cluster_system_node_group_node_count" {
  type = number
}

variable "aws_cluster_system_node_group_instance_types" {
  type = list(string)
}

variable "aws_cluster_service_ipv4_cidr" {
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
