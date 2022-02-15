terraform {
  required_version = "=1.1.5"

  required_providers {
    aws = {
      version = "~> 4.1.0"
    }

    azurerm = {
      version = "~> 2.96.0"
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

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
