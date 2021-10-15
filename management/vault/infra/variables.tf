variable "management_infra_location" {
  type = string
}

variable "management_infra_aks_name" {
  type = string
}

variable "management_infra_key_vault_name" {
  type = string
}

variable "management_infra_domain" {
  type = string
}

variable "management_vault_infra_storage_account_name" {
  type = string
}

variable "management_vault_infra_image_repository" {
  type = string
}

variable "management_vault_infra_image_tag" {
  type = string
}

variable "management_vault_infra_chart_version" {
  type = string
}

variable "management_vault_infra_disable_affinity" {
  type    = bool
  default = false
}

variable "azuread_client_id" {
  type    = string
  default = ""
}

variable "azuread_client_secret" {
  type    = string
  default = ""
}
