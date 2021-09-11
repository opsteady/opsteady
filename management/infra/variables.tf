variable "management_infra_location" {
  type = string
}

variable "management_infra_acr_name" {
  type = string
}

variable "management_infra_log_analytics_workspace_retention" {
  type = number
}

variable "management_infra_platform_admins" {
  type    = list(string)
  default = []
}

variable "management_infra_platform_admin_owners" {
  type    = list(string)
  default = []
}
variable "management_infra_platform_developers" {
  type    = list(string)
  default = []
}
variable "management_infra_platform_developer_owners" {
  type    = list(string)
  default = []
}

variable "management_infra_platform_viewers" {
  type    = list(string)
  default = []
}
variable "management_infra_platform_viewer_owners" {
  type    = list(string)
  default = []
}

variable "management_infra_vnet_address_space" {
  type = list(any)
}
variable "management_infra_azure_subnet_pods_address_prefixes" {
  type = list(any)
}

variable "management_infra_azure_subnet_public_address_prefixes" {
  type = list(any)
}

variable "management_infra_domain" {
  type = string
}
