terraform {
  required_version = "=1.1.3"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.90.0"
    }
  }

  /*
   * Comment this backend when bootstrapping for the first time. This
   * will result in a local statefile. After a successful local bootstrap,
   * uncomment this backend and run Terraform init/apply again. This will
   * result in Terraform asking to push the local state to the remote location.
   * After this is done, remove the local statefile and its backup,
   * commit the remote backend and work from the remote state.
   */
  backend "azurerm" {
    resource_group_name  = "terraform-state"
    container_name       = "management"
    key                  = "azure/management/management-bootstrap.tfstate"
  }
}

provider "azurerm" {
  features {}
}
