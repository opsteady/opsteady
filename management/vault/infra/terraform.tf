terraform {
  required_version = "=1.0.9"

  required_providers {
    azurerm = {
      version = "~> 2.83.0"
    }

    azuread = {
      version = "~> 2.8.0"
    }

    kubernetes = {
      version = "~> 2.6.0"
    }

    helm = {
      version = "~> 2.3.0"
    }

    tls = {
      version = "~> 3.1.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "terraform-state"
    container_name       = "management"
    key                  = "azure/management/management-vault-infra.tfstate"
  }
}

provider "azurerm" {
  features {}
}

provider "azuread" {
  client_id     = var.azuread_client_id
  client_secret = var.azuread_client_secret
}

provider "kubernetes" {
  host                   = data.azurerm_kubernetes_cluster.management.kube_config.0.host
  username               = data.azurerm_kubernetes_cluster.management.kube_config.0.username
  password               = data.azurerm_kubernetes_cluster.management.kube_config.0.password
  client_key             = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.client_key)
  client_certificate     = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.client_certificate)
  cluster_ca_certificate = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.cluster_ca_certificate)
}

provider "helm" {
  kubernetes {
    host                   = data.azurerm_kubernetes_cluster.management.kube_config.0.host
    username               = data.azurerm_kubernetes_cluster.management.kube_config.0.username
    password               = data.azurerm_kubernetes_cluster.management.kube_config.0.password
    client_key             = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.client_key)
    client_certificate     = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.client_certificate)
    cluster_ca_certificate = base64decode(data.azurerm_kubernetes_cluster.management.kube_config.0.cluster_ca_certificate)
  }
}
