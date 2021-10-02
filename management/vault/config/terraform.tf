terraform {
  required_version = "= 1.0.5"

  required_providers {
    azurerm = {
      version = "~> 2.76.0"
    }

    vault = {
      source  = "hashicorp/vault"
      version = "~> 2.24.0"
    }

    azuread = {
      version = "~> 2.2.1"
    }

    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.59.0"
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
  address      = "https://vault.management.${var.management_infra_domain}"
  token        = var.management_vault_config_token
  ca_cert_file = var.management_vault_config_ca_cert_file
}

provider "aws" {
  region = "eu-west-1"
}
