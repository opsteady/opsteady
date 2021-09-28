variable "name" {
  type = string
}

variable "required_resource_access" {
  type = list(object({
    resource_app_id = string
    resource_access = list(object({
      id   = string
      type = string
    }))
  }))
  default = []
}

variable "app_roles" {
  type = list(object({
    id                   = string
    allowed_member_types = list(string)
    description          = string
    display_name         = string
    is_enabled           = bool
    value                = string
  }))
  default = []
}

variable "owners" {
  type    = list(string)
  default = []
}

variable "redirect_uris" {
  type    = list(string)
  default = []
}

variable "group_membership_claims" {
  type    = list(string)
  default = null
}
