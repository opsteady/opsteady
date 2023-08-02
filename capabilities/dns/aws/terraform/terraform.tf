terraform {
  required_version = "=1.3.6"

  required_providers {
    aws = {
      version = "~> 4.48.0"
    }

    azurerm = {
      version = "~> 3.37.0"
    }

    vault = {
      version = "~> 3.19.0"
    }

    tls = {
      version = "~> 4.0.0"
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

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
