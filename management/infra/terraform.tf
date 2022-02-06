terraform {
  required_version = "=1.1.5"

  required_providers {
    azurerm = {
      version = "~> 2.95.0"
    }

    azuread = {
      version = "~> 2.17.0"
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
  host                   = var.management_infra_minimal ? "" : azurerm_kubernetes_cluster.management.0.kube_config.0.host
  username               = var.management_infra_minimal ? "" : azurerm_kubernetes_cluster.management.0.kube_config.0.username
  password               = var.management_infra_minimal ? "" : azurerm_kubernetes_cluster.management.0.kube_config.0.password
  client_key             = var.management_infra_minimal ? "" : base64decode(azurerm_kubernetes_cluster.management.0.kube_config.0.client_key)
  client_certificate     = var.management_infra_minimal ? "" : base64decode(azurerm_kubernetes_cluster.management.0.kube_config.0.client_certificate)
  cluster_ca_certificate = var.management_infra_minimal ? "" : base64decode(azurerm_kubernetes_cluster.management.0.kube_config.0.cluster_ca_certificate)
}
