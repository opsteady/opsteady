terraform {
  required_version = "= 1.0.8"

  required_providers {
    azurerm = {
      version = "~> 2.76.0"
    }

    azuread = {
      version = "~> 2.5.0"
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
}

provider "azurerm" {
  subscription_id = var.foundation_azure_subscription_id
  features {}
}

provider "azurerm" {
  alias = "management"

  client_id       = var.management_client_id
  client_secret   = var.management_client_secret
  subscription_id = var.management_subscription_id
  tenant_id       = var.tenant_id
  features {}
}
