terraform {
  required_version = "=1.3.6"

  required_providers {
    azurerm = {
      version = "~> 3.37.0"
    }

    vault = {
      source  = "hashicorp/vault"
      version = "~> 3.11.0"
    }

    aws = {
      version = "~> 4.48.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "aws" {
  region = var.aws_foundation_region
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
