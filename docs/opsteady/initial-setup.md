# Initial setup of the Management platform

This documentation is used to initially setup the management platform as a way to bootstrap everything and to be able to use CI/CD from then on.

## 00 Requirements

We assume:

- you have access to Azure
- you have created a **management** subscription
- you have admin rights on the **management** subscription
- you have Docker installed

Note: some names used below need to be adjusted as they are globally unique in Azure

## 01 Docker images

You can do the initial setup from your local machine using your own tools but as mentioned in the [ADR](../adr/0011-no-local-tools.md) we want this to be executable without local tools. For that you can create the docker image to do that.

```bash
cd docker/base
docker build -t dev-management.azurecr.io/base:1.0.0 .
cd ../cicd
docker build --build-arg ACR_NAME=dev-management -t dev-management.azurecr.io/cicd:1.0.0 .
cd ../..
```

## 02 Bootstrap

Either open the project in Visual Studio code with the provided remote container or use the container manually:

```bash
export ACR_NAME=dev-management
docker run -it --rm -v $(pwd):/data dev-management.azurecr.io/cicd:1.0.0 /bin/bash
```

Actual steps to perform in the container locally or VSC remote container or on your local machine

```bash
az login --use-device-code
cd management/bootstrap
az account set --subscription management
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init
terraform plan -var="management_bootstrap_terraform_state_location=westeurope" -var="management_bootstrap_terraform_state_account_name=devmgmweu"
terraform apply -var="management_bootstrap_terraform_state_location=westeurope" -var="management_bootstrap_terraform_state_account_name=devmgmweu"
```

Uncomment the lines behind `backend "azurerm"` in the `provider.tf` to upload the state to the remote backend.

```bash
terraform init -reconfigure -backend-config "storage_account_name=devmgmweu"
```

`terraform.tfstate` will be empty and as it is uploaded to the remote storage, it can be safely deleted together with `terraform.tfstate.backup`.
