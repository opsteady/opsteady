terraform {
  required_version = "=1.1.3"

  required_providers {
    azurerm = {
      version = "~> 2.92.0"
    }

    vault = {
      version = "~> 2.24.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "azurerm" {
  subscription_id = var.foundation_azure_subscription_id
  features {}
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
