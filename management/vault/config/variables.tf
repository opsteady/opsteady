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

# A list of AWS accounts
variable "management_vault_config_accounts" {
  type = list(object({
    name       = string
    id         = string
    access_key = string
    secret_key = string
  }))
}

variable "azuread_client_id" {
  type    = string
  default = ""
}

variable "azuread_client_secret" {
  type    = string
  default = ""
}
