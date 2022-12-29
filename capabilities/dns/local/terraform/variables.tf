variable "local_foundation_name" {
  type = string
}

variable "local_foundation_location" {
  type = string
}

variable "local_foundation_resource_group" {
  type = string
}

variable "local_foundation_subscription_id" {
  type = string
}

variable "local_foundation_public_zone_id" {
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

variable "azuread_client_id" {
  type    = string
  default = ""
}

variable "azuread_client_secret" {
  type    = string
  default = ""
}
