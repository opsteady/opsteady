terraform {
  required_version = "=1.0.11"

  required_providers {
    aws = {
      version = "~> 3.65.0"
    }

    azurerm = {
      version = "~> 2.86.0"
    }

    vault = {
      version = "~> 2.24.0"
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

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
