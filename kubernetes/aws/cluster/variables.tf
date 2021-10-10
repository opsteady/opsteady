variable "foundation_aws_region" {
  type = string
}

variable "foundation_aws_name" {
  type = string
}

variable "foundation_aws_vpc_id" {
  type = string
}

variable "foundation_aws_eks_a_subnet_id" {
  type = string
}

variable "foundation_aws_eks_b_subnet_id" {
  type = string
}

variable "foundation_aws_eks_c_subnet_id" {
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

variable "kubernetes_aws_cluster_public_access_cidrs" {
  type = list(string)
}

variable "kubernetes_aws_cluster_kubernetes_version" {
  type = string
}

variable "kubernetes_aws_cluster_system_node_group_node_count" {
  type = number
}

variable "kubernetes_aws_cluster_system_node_group_instance_types" {
  type = list(string)
}

variable "kubernetes_aws_cluster_service_ipv4_cidr" {
  type = string
}
