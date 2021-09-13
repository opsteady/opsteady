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
# Run this from the root of the repository
docker run -it --rm -v $(pwd):/data dev-management.azurecr.io/cicd:1.0.0 /bin/bash
```

Before you start, comment the entire `backend "azurerm"` section in `management/bootstrap/terraform.tf`. This will keep the state local temporarily.

When bootstrapping we provide a set of defaults for the Terraform values. They are located in `management/defaults/bootstrap.default.tfvars`. You can choose to use this file as-is, or copy it to `management/defaults/bootstrap.tfvars` and adjust the values. This custom file will be ignored in the Git repository and is for your use only.

Actual steps to perform in the container locally, VSC remote container or on your local machine:

```bash
az login --use-device-code
az account set --subscription management
cd management/bootstrap
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init

terraform apply -var-file=../defaults/bootstrap.default.tfvars   # Adjust to ../defaults/bootstrap.tfvars if you have a custom variables file
```

Uncomment the entire `backend "azurerm"` section in `management/bootstrap/terraform.tf` to upload the state to the remote backend. **Make sure to update the `storage_account_name` setting to the name that you used in the previous apply step.**:

```bash
terraform init -reconfigure -backend-config "storage_account_name=This name should match management_bootstrap_terraform_state_account_name"
```

`terraform.tfstate` will be empty and as it is uploaded to the remote storage, it can be safely deleted together with `terraform.tfstate.backup`.

## 03 Infra

Update update the `storage_account_name` setting in `management/infra/terraform.tf` to the name that you used in the previous bootstrap.

When creating the initial management infra, we provide a set of defaults for the Terraform values. They are located in `management/defaults/infra.default.tfvars`. To deploy the management infrastructure successfully, you will have to copy the default file to a custom tfvars file called `management/defaults/infra.tfvars`. In this file you need to update the `management_infra_key_vault_ip_rules` to include your IP address in the (now) empty list. This makes sure that when Terraform needs to configure the Key Vault it can do so from your location.
`

```bash
cd management/infra
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init
terraform apply -var-file=../defaults/infra.tfvars
```

After a successful apply, you can delete the custom `management/defaults/infra.tfvars` file because you will not need it anymore.
