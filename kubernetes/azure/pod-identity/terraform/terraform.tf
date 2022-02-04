terraform {
  required_version = "=1.1.4"

  required_providers {
    azurerm = {
      version = "~> 2.95.0"
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
