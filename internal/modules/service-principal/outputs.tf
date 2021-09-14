output "azuread_application_application_id" {
  value = azuread_application.spn.application_id
}

output "azuread_application_object_id" {
  value = azuread_application.spn.object_id
}

output "azuread_service_principal_application_id" {
  value = azuread_service_principal.spn.application_id
}

output "azuread_service_principal_object_id" {
  value = azuread_service_principal.spn.object_id
}

output "azuread_service_principal_password" {
  value     = azuread_service_principal_password.spn.value
  sensitive = true
}
