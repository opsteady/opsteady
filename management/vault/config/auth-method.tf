module "vault_oidc_auth" {
  source                  = "../../../internal/modules/service-principal"
  name                    = "vault-oidc-auth"
  group_membership_claims = ["All"]
  redirect_uris = [
    "https://vault.management.${var.management_infra_domain}/ui/vault/auth/oidc/oidc/callback",
    "http://localhost:8250/oidc/callback",
  ]

  required_resource_access = [
    {
      resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

      resource_access = [{
        id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["GroupMember.Read.All"]
        type = "Scope"
      }]
    }
  ]
}

resource "vault_jwt_auth_backend" "management_azure_ad" {
  description        = "Management Azure AD login"
  path               = "oidc"
  type               = "oidc"
  oidc_discovery_url = "https://login.microsoftonline.com/${data.azurerm_client_config.current.tenant_id}/v2.0"
  oidc_client_id     = module.vault_oidc_auth.azuread_application_application_id
  oidc_client_secret = module.vault_oidc_auth.azuread_service_principal_password
  default_role       = "platform-admin"

  tune {
    listing_visibility = "unauth"
    default_lease_ttl  = "10h"
    max_lease_ttl      = "24h"
    token_type         = "default-service"
  }
}

resource "vault_jwt_auth_backend_role" "platform_viewer" {
  backend        = vault_jwt_auth_backend.management_azure_ad.path
  role_name      = "platform-reader"
  token_policies = ["platform-reader"]

  oidc_scopes = ["https://graph.microsoft.com/.default"]
  user_claim  = "email"
  role_type   = "oidc"
  allowed_redirect_uris = [
    "https://vault.management.${var.management_infra_domain}/ui/vault/auth/oidc/oidc/callback",
  ]
  groups_claim = "groups"
  bound_claims = { "groups" = data.azuread_group.platform_viewer.id }
}

resource "vault_jwt_auth_backend_role" "platform_operator" {
  backend        = vault_jwt_auth_backend.management_azure_ad.path
  role_name      = "platform-operator"
  token_policies = ["platform-operator"]

  oidc_scopes = ["https://graph.microsoft.com/.default"]
  user_claim  = "email"
  role_type   = "oidc"
  allowed_redirect_uris = [
    "https://vault.management.${var.management_infra_domain}/ui/vault/auth/oidc/oidc/callback",
  ]
  groups_claim = "groups"
  bound_claims = { "groups" = data.azuread_group.platform_operator.id }
}

resource "vault_jwt_auth_backend_role" "platform_admin" {
  backend        = vault_jwt_auth_backend.management_azure_ad.path
  role_name      = "platform-admin"
  token_policies = ["platform-admin"]

  oidc_scopes = ["https://graph.microsoft.com/.default"]
  user_claim  = "email"
  role_type   = "oidc"
  allowed_redirect_uris = [
    "https://vault.management.${var.management_infra_domain}/ui/vault/auth/oidc/oidc/callback",
  ]
  groups_claim = "groups"
  bound_claims = { "groups" = data.azuread_group.platform_admin.id }
}
