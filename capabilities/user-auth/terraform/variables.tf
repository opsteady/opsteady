variable "aws_foundation_name" {
  type    = string
  default = ""
}

variable "azure_foundation_name" {
  type    = string
  default = ""
}

variable "local_foundation_name" {
  type    = string
  default = ""
}

variable "capabilities_user_auth_oidc_owners" {
  type = list(string)
}

variable "azure_foundation_public_zone_name" {
  type    = string
  default = ""
}

variable "aws_foundation_public_zone_name" {
  type    = string
  default = ""
}

variable "local_foundation_public_zone_name" {
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
