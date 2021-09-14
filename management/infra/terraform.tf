terraform {
  required_version = "= 1.0.5"

  required_providers {
    azurerm = {
      version = "~> 2.76.0"
    }

    azuread = {
      version = "~> 2.2.1"
    }

    kubernetes = {
      version = "~> 2.4.1"
    }

    helm = {
      version = "~> 2.3.0"
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
    storage_account_name = "This name should match management_bootstrap_terraform_state_account_name"
    container_name       = "management"
    key                  = "infra.tfstate"
  }
}

provider "azurerm" {
  features {}
}

provider "azuread" {}

provider "kubernetes" {
  host                   = azurerm_kubernetes_cluster.management.kube_config.0.host
  username               = azurerm_kubernetes_cluster.management.kube_config.0.username
  password               = azurerm_kubernetes_cluster.management.kube_config.0.password
  client_key             = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_key)
  client_certificate     = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_certificate)
  cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.cluster_ca_certificate)
}

provider "helm" {
  kubernetes {
    host                   = azurerm_kubernetes_cluster.management.kube_config.0.host
    username               = azurerm_kubernetes_cluster.management.kube_config.0.username
    password               = azurerm_kubernetes_cluster.management.kube_config.0.password
    client_key             = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_key)
    client_certificate     = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.client_certificate)
    cluster_ca_certificate = base64decode(azurerm_kubernetes_cluster.management.kube_config.0.cluster_ca_certificate)
  }
}
