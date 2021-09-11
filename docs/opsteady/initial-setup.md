# Initial setup of the Management platform

This documentation is used to initially setup the management platform as a way to bootstrap everything and to be able to use CI/CD from then on. It builds based on [ADR management setup](../adr/0013-management-setup.md)

## 00 Requirements

We assume:

- you have access to Azure
- you have created a **management** subscription
- you have admin rights on the **management** subscription
- you have a (sub)domain that you can use
- you have Docker installed

Note-1: The (sub)domain like os.opsteady.com is used create subdomains for every platform, for example management.os.opsteady.com or dev-aws.os.opsteady.com. On top of these subdomains we will expose applications, like vault.management.os.opsteady.com.

Note-2: some names used below need to be adjusted as they are globally unique in Azure

## 01 Docker images

You can do the initial setup from your local machine using your own tools but as mentioned in the [ADR](../adr/0011-no-local-tools.md) we want this to be executable without local tools. For that you can create the docker image.

```bash
cd docker/base
docker build -t dev-management.azurecr.io/base:1.0.0 .
cd ../cicd
docker build --build-arg ACR_NAME=dev-management -t dev-management.azurecr.io/cicd:1.0.0 .
cd ../..
```

Note-1: Make sure to replace `dev-management.azurecr.io` with whatever container registry you are using.

## 02 Bootstrap

Either open the project in Visual Studio code with the provided remote container or use the container manually:

```bash
export ACR_NAME=dev-management
docker run -it --rm -v $(pwd):/data dev-management.azurecr.io/cicd:1.0.0 /bin/bash
```

Before you start, comment the entire `backend "azurerm"` section in `terraform.tf`. This will keep the state local temporarily.

Actual steps to perform in the container locally, VSC remote container or on your local machine:

```bash
az login --use-device-code
az account set --subscription management
cd management/bootstrap
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init
terraform plan -var="management_bootstrap_terraform_state_location=westeurope" -var="management_bootstrap_terraform_state_account_name=devmgmtweu"
terraform apply -var="management_bootstrap_terraform_state_location=westeurope" -var="management_bootstrap_terraform_state_account_name=devmgmtweu"
```

Uncomment the lines behind `backend "azurerm"` in `terraform.tf` to upload the state to the remote backend:

```bash
terraform init -reconfigure -backend-config "storage_account_name=devmgmtweu"
```

`terraform.tfstate` will be empty and as it is uploaded to the remote storage, it can be safely deleted together with `terraform.tfstate.backup`.

## 03 Infra

```bash
cd management/infra
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init -backend-config "storage_account_name=devmgmweu"
terraform plan \
  -var='management_infra_acr_name=devmgmtweu' \
  -var='management_infra_vnet_address_space=["10.0.0.0/19"]'\
  -var='management_infra_azure_subnet_pods_address_prefixes=["10.0.0.0/20"]' \
  -var='management_infra_platform_admins=[]' \
  -var='management_infra_platform_admin_owners=[]'\
  -var='management_infra_platform_developers=[]' \
  -var='management_infra_platform_developer_owners=[]' \
  -var='management_infra_platform_viewers=[]' \
  -var='management_infra_platform_viewer_owners=[]' \
  -var='management_infra_domain=os.opsteady.com' \
  -var='management_infra_location=westeurope' \
  -var='management_infra_log_analytics_workspace_retention=7' \
  -var='management_infra_azure_subnet_public_address_prefixes=["10.0.16.0/24"]
```
