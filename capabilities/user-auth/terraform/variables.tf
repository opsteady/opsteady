variable "foundation_aws_name" {
  type    = string
  default = ""
}

variable "foundation_azure_name" {
  type    = string
  default = ""
}
variable "capabilities_user_auth_oidc_owners" {
  type = list(string)
}

variable "foundation_azure_public_zone_name" {
  type    = string
  default = ""
}

variable "foundation_aws_public_zone_name" {
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

variable "platform_version" {
  type = string
}

variable "platform_environment_name" {
  type = string
}

variable "platform_component_name" {
  type = string
}
