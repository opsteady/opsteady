terraform {
  required_version = "=1.1.6"

  required_providers {
    aws = {
      version = "~> 4.1.0"
    }

    azurerm = {
      version = "~> 2.99.0"
    }

    vault = {
      version = "~> 3.2.0"
    }

    tls = {
      version = "~> 3.1.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "aws" {
  region = var.foundation_aws_region
}

provider "azurerm" {
  alias = "management"

  client_id       = var.management_client_id
  client_secret   = var.management_client_secret
  subscription_id = var.management_subscription_id
  tenant_id       = var.tenant_id
  features {}
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
