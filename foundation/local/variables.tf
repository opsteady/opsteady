variable "local_foundation_name" {
  description = "Name to be used for resources or as a suffix, mostly plt1"
  type        = string
}

variable "local_foundation_public_name" {
  description = "The name used as the sub domain"
  type        = string
}

variable "local_foundation_subscription_id" {
  type = string
}

variable "local_foundation_location" {
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

variable "platform_terraform_output_path" {
  type = string
}

variable "platform_vault_vars_name" {
  type = string
}
