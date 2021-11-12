terraform {
  required_version = "=1.0.11"

  required_providers {
    azurerm = {
      version = "~> 2.85.0"
    }

    vault = {
      source  = "hashicorp/vault"
      version = "~> 2.24.0"
    }

    aws = {
      version = "~> 3.65.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "terraform-state"
    storage_account_name = "aiplatformmgmt"
    container_name       = "platform"
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
