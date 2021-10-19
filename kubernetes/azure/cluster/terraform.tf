terraform {
  required_version = "= 1.0.9"

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
  subscription_id = var.foundation_azure_subscription_id
  features {}
}
