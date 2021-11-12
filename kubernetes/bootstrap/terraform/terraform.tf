terraform {
  required_version = "=1.0.11"

  required_providers {
    azuread = {
      version = "~> 2.9.0"
    }

    kubernetes = {
      version = "~> 2.6.0"
    }

    vault = {
      version = "~> 2.24.0"
    }

    time = {
      version = "~> 0.7.2"
    }
  }

  backend "azurerm" {
    resource_group_name = "terraform-state"
    container_name      = "platform"
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
  tenant_id     = var.azuread_tenant_id
}

provider "vault" {
  address = var.vault_address
  token   = var.vault_token
}
