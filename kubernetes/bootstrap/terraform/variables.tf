// This is variable is set for AWS platforms
variable "aws_foundation_environment_name" {
  type    = string
  default = ""
}

// This is variable is set for Azure platforms
variable "azure_foundation_environment_name" {
  type    = string
  default = ""
}

// This is variable is set for Local platforms
variable "local_foundation_environment_name" {
  type    = string
  default = ""
}
variable "azuread_tenant_id" {
  type = string
}

variable "azuread_client_id" {
  type = string
}

variable "azuread_client_secret" {
  type = string
}

variable "management_infra_acr_name" {
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
