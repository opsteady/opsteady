# This service principal represents the Microsoft Graph service. It is used to find the identifier of Graph permissions.
# This will replace the below legacy AAD Graph API in the future.
resource "azuread_service_principal" "msgraph" {
  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing   = true
}

# This service principal represents the Active Directory Graph service. It is used to find the identifier of Graph permissions.
# This will be replaced by the above Microsoft Graph API in the future but currently this is the only mechanism that Vault
# supports to generate Azure credentials.
resource "azuread_service_principal" "aadgraph" {
  application_id = data.azuread_application_published_app_ids.well_known.result.AzureActiveDirectoryGraph
  use_existing   = true
}
