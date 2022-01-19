terraform {
  required_version = "=1.1.3"

  required_providers {
    vault = {
      version = "~> 3.2.0"
    }

    azuread = {
      version = "~> 2.15.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
  tenant_id     = var.azuread_tenant_id
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
