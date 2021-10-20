// This is variable is set for AWS platforms
variable "foundation_aws_environment_name" {
  type = string
  default = ""
}

// This is variable is set for Azure platforms
variable "foundation_azure_environment_name" {
  type = string
  default = ""
}

variable "azuread_tenant_id" {
  type    = string
}

variable "azuread_client_id" {
  type    = string
}

variable "azuread_client_secret" {
  type    = string
}

variable "management_infra_acr_name" {
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
