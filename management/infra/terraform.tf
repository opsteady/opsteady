terraform {
  required_version = "= 1.0.5"

  required_providers {
    azurerm = {
      version = "~> 2.74.0"
    }

    azuread = {
      version = "~> 2.1.0"
    }
  }

  # backend "azurerm" {
  #   resource_group_name  = "terraform-state"
  #   storage_account_name = "This name should match management_bootstrap_terraform_state_account_name"
  #   container_name       = "management"
  #   key                  = "infra.tfstate"
  # }
}

provider "azurerm" {
  features {}
}

provider "azuread" {}
