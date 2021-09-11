terraform {
  required_version = "= 1.0.5"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 2.74.0"
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
    storage_account_name = "This name should match management_bootstrap_terraform_state_account_name"
    container_name       = "management"
    key                  = "bootstrap.tfstate"
  }
}

provider "azurerm" {
  features {}
}