variable "foundation_local_name" {
  description = "Name to be used for resources or as a suffix, mostly plt1"
  type        = string
}

variable "foundation_local_public_name" {
  description = "The name used as the sub domain"
  type        = string
}

variable "foundation_local_subscription_id" {
  type = string
}

variable "foundation_local_location" {
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
