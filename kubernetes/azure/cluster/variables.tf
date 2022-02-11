variable "azure_cluster_system_nodepool_node_count" {
  type = number
}

variable "azure_cluster_system_nodepool_node_size" {
  type = string
}

variable "azure_cluster_platform_nodepool_node_count" {
  type = number
}

variable "azure_cluster_platform_nodepool_node_size" {
  type = string
}

variable "azure_cluster_kubernetes_version" {
  type = string
}

variable "azure_foundation_subscription_id" {
  type = string
}

variable "azure_foundation_name" {
  type = string
}

variable "azure_foundation_location" {
  type = string
}

variable "foundation_azure_key_vault_id" {
  type = string
}

variable "foundation_azure_public_ip_prefix_id" {
  type = string
}

variable "foundation_azure_pods_subnet_id" {
  type = string
}
variable "foundation_azure_log_analytics_id" {
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
