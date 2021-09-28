resource "vault_policy" "platform_viewer" {
  name   = "platform-viewer"
  policy = file("policies/platform-viewer.hcl")
}

resource "vault_policy" "platform_operator" {
  name   = "platform_operator"
  policy = file("policies/platform-operator.hcl")
}

resource "vault_policy" "platform_admin" {
  name   = "platform_admin"
  policy = file("policies/platform-admin.hcl")
}
