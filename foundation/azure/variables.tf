variable "foundation_azure_subscription_id" {
  type = string
}

variable "foundation_azure_location" {
  type = string
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

variable "azuread_client_id" {
  type = string
}

variable "azuread_client_secret" {
  type = string
}

variable "foundation_azure_name" {
  type = string
}

variable "foundation_azure_log_analytics_workspace_retention" {
  type = number
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

variable "tenant_id" {
  type = string
}

variable "management_infra_key_vault_administrators" {
  type = list(any)
}
