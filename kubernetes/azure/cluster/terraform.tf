terraform {
  required_version = "=1.3.6"

  required_providers {
    azurerm = {
      version = "~> 2.78.0" // NOTE: Don't update until this is fixed: https://github.com/Azure/AKS/issues/2584
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "azurerm" {
  subscription_id = var.azure_foundation_subscription_id
  features {}
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
