terraform {
  required_version = "=1.3.6"

  required_providers {
    azurerm = {
      version = "~> 3.2.0"
    }

    vault = {
      version = "~> 3.11.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "azurerm" {
  subscription_id = var.local_foundation_subscription_id
  features {}
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
}
