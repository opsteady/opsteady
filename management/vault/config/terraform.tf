terraform {
  required_version = "=1.0.9"

  required_providers {
    azurerm = {
      version = "~> 2.83.0"
    }

    vault = {
      source  = "hashicorp/vault"
      version = "~> 2.24.0"
    }

    azuread = {
      version = "~> 2.8.0"
    }

    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.63.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "management"
    key                 = "azure/management/management-vault-config.tfstate"
  }
}

provider "azurerm" {
  features {}
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
}

provider "vault" {
  address = "https://vault.management.${var.management_infra_domain}"
  token   = var.vault_token
}

provider "aws" {
  region = "eu-west-1"
}
