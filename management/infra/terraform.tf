terraform {
  required_version = "=1.1.3"

  required_providers {
    azurerm = {
      version = "~> 2.92.0"
    }

    azuread = {
      version = "~> 2.15.0"
    }

    kubernetes = {
      version = "~> 2.7.0"
    }

    dns = {
      version = "~> 3.2.1"
    }

    time = {
      version = "~> 0.7.2"
    }

    tls = {
      version = "~> 3.1.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "terraform-state"
    container_name       = "management"
    key                  = "azure/management/management-infra.tfstate"
  }
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
}

provider "azurerm" {
  features {}
}

provider "kubernetes" {
  host                   = azurerm_kubernetes_cluster.management.kube_config.0.host
  username               = azurerm_kubernetes_cluster.management.kube_config.0.username
  password               = azurerm_kubernetes_cluster.management.kube_config.0.password
  client_key             = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_key)
  client_certificate     = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_certificate)
  cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.cluster_ca_certificate)
}
