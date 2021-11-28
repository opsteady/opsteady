resource "azuread_application" "oidc" {
  display_name = "oidc-${local.foundation_name}"

  owners = var.capabilities_user_auth_oidc_owners

  group_membership_claims        = ["All"]
  fallback_public_client_enabled = false

  web {
    redirect_uris = [
      "https://${local.oidc_callback_url}",
    ]
  }
}

resource "azuread_service_principal" "oidc" {
  application_id               = azuread_application.oidc.application_id
  app_role_assignment_required = false
}


resource "azuread_service_principal_password" "oidc" {
  service_principal_id = azuread_service_principal.oidc.id
}
