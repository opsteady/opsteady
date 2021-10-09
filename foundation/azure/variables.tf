variable "foundation_azure_name" {
  description = "Name to be used for resources or as a suffix, mostly plt1"
  type        = string
}

variable "foundation_azure_environment_name" {
  description = "Name of the platform environment, for example dev-azure"
  type        = string
}

variable "foundation_azure_subscription_id" {
  type = string
}

variable "foundation_azure_public_name" {
  description = "The name used as the sub domain"
  type        = string
}

variable "foundation_azure_location" {
  type = string
}

variable "foundation_azure_log_analytics_workspace_retention" {
  description = "It has to be more than 30"
  type        = number
}

variable "foundation_azure_vnet_address_space" {
  type = list(any)
}

variable "foundation_azure_subnet_public_address_prefixes" {
  type = list(any)
}

variable "foundation_azure_subnet_pods_address_prefixes" {
  type = list(any)
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

# Used for creating output to Vault
variable "vault_address" {
  type = string
}

variable "vault_token" {
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
