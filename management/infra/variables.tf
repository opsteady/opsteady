variable "management_infra_location" {
  type = string
}

variable "management_infra_acr_name" {
  type = string
}

variable "management_infra_aks_name" {
  type = string
}

variable "management_infra_aks_sku_tier" {
  type = string
}

variable "management_infra_aks_kubernetes_version" {
  type = string
}

variable "management_infra_aks_system_node_count" {
  type = number
}

variable "management_infra_aks_system_node_size" {
  type = string
}

variable "management_infra_aks_api_server_authorized_ip_ranges" {
  type = list(any)
}

variable "management_infra_aks_system_nodepool_node_count" {
  type = number
}

variable "management_infra_aks_system_nodepool_node_size" {
  type = string
}

variable "management_infra_log_analytics_workspace_retention" {
  type = number
}

variable "management_infra_key_vault_name" {
  type = string
}

variable "management_infra_key_vault_ip_rules" {
  type = list
}

variable "management_infra_key_vault_administrators" {
  description = "A list of principal IDs that can administer the key vault"
  type = list
}

variable "management_infra_platform_admins" {
  type    = list(string)
  default = []
}

variable "management_infra_platform_admin_owners" {
  type = list(string)
}
variable "management_infra_platform_developers" {
  type = list(string)
}
variable "management_infra_platform_developer_owners" {
  type = list(string)
}

variable "management_infra_platform_viewers" {
  type = list(string)
}
variable "management_infra_platform_viewer_owners" {
  type = list(string)
}

variable "management_infra_vnet_address_space" {
  type = list(any)
}
variable "management_infra_azure_subnet_pods_address_prefixes" {
  type = list(any)
}

variable "management_infra_azure_subnet_public_address_prefixes" {
  type = list(any)
}

variable "management_infra_domain" {
  type = string
}
