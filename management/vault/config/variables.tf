variable "management_infra_domain" {
  type = string
}

variable "vault_token" {
  type = string
}

variable "management_vault_config_ca_cert_file" {
  type = string
}

variable "management_vault_config_aws_region" {
  type = string
}

# A map of Azure subscription IDs and names.
variable "management_vault_config_subscriptions" {
  type = map(string)
}

# A map of AWS account IDs and names.
variable "management_vault_config_accounts" {
  type = map(string)
}

variable "azuread_client_id" {
  type = string
  default = ""
}

variable "azuread_client_secret" {
  type = string
  default = ""
}
