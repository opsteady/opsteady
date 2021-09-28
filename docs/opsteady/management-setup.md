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

Update the `storage_account_name` setting in `management/infra/terraform.tf` to the name that you used in the previous bootstrap.

When creating the initial management infra, we provide a set of defaults for the Terraform values. They are located in `management/defaults/infra.default.tfvars`. To deploy the management infrastructure successfully, copy the default file to a custom tfvars file called `management/defaults/infra.tfvars`. In this file update the `management_infra_key_vault_ip_rules` to include your IP address in the (now) empty list. This makes sure that when Terraform needs to configure the Key Vault it can do so from your location.
`

```bash
cd management/infra
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"
terraform apply -var-file=../defaults/infra.tfvars

# Run the Terraform plan again to make sure that everything is applied.
# The virtual network rules in the Key Vault network ACL might not apply
# the first time for some unknown reason.
terraform apply -var-file=../defaults/infra.tfvars
```

After a successful apply, you can delete the custom `management/defaults/infra.tfvars` file because you will not need it anymore.

At this point a DNS zone has been created for the Opsteady platform. From your DNS hosting provider you need to delegate a subzone to this domain.

```bash
terraform state show azurerm_dns_zone.public_root
```

The output of this command will show you the name servers (amongst other things) that you need to delegate to. The `management_infra_domain` variable in the infra defaults contains the subdomain that you need to delegate to. Create the NS records with your DNS hosting provider. It can take some time before the DNS resolving is active.

## 04 Vault Infrastructure

When creating the Vault infrastructure, we provide a set of defaults for the Terraform values. They are located in `management/defaults/vault-infra.default.tfvars`. To deploy the Vault infrastructure successfully, copy the default file to a custom tfvars file called `management/defaults/vault-infra.tfvars`. In this file update the `management_vault_infa_storage_account_name` to a unique name. This storage account will host the Vault CA certificate. Make sure that the `management_infra_*` settings are equal to the values that you used in the management infra step. Vault builds on top of the management infra and needs to locate certain resources based on these names.

```bash
cd management/vault/infra
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"
terraform apply -var-file=../../defaults/vault-infra.tfvars
```

Vault is now deployed to the cluster but all the instances are in a sealed state. We need initialise the cluster with the following steps:

```bash
az aks get-credentials -g management -n management --admin
kubectl exec -n platform -it vault-0 -- vault operator init -ca-path=/vault/userconfig/vault-tls/ca.crt
```

**If the command succeeds you will see the recovery keys and the initial root token for Vault. Store this in a secure location and distribute the recovery keys to trusted parties.**

The certificate authority file for Vault can be downloaded from `https://${management_vault_infra_storage_account_name}.blob.core.windows.net/vault-ca/ca.pem`. With this file you should be able to connect securely to Vault on `https://vault.management.${management_infra_domain}`.

## 04 Vault Configuration

When creating the Vault configuration, we provide a set of defaults for the Terraform values. They are located in `management/defaults/vault-config.default.tfvars`. To deploy the Vault configuration successfully, copy the default file to a custom tfvars file called `management/defaults/vault-config.tfvars`. Add the management subscription ID as a value to the `management_vault_config_subscriptions` map. Make sure that the `management_infra_*` settings are equal to the values that you used in the management infra step. Vault builds on top of the management infra and needs to locate certain resources based on these names.

The `management_vault_config_subscriptions` variable will contain all the subscriptions that are going to be managed by the Opsteady platform. The `management_vault_config_accounts` variable will contain all the AWS accounts that are managed by the Opsteady platform. As the platform grows, these variables will be extended. For now it's just the management subscriptions that we want to bring under management.

When doing the initial Vault configuration (when the management subscription is not yet managed by Opsteady CI/CD) make sure to comment the `client_id` and `cient_secret` settings in the azuread provider in the `terraform.tf` file.

```bash
cd management/vault/config
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"

# Before we can apply the Terraform code, we need to grab the Vault CA certificate and put it in a well-known location, so that the Vault provider can use it.
curl -o vault-ca.pem https://${management_vault_infra_storage_account_name}.blob.core.windows.net/vault-ca/ca.pem

terraform apply -var-file=../../defaults/vault-config.tfvars -var management_vault_config_token=$VAULT_ROOT_TOKEN_FROM_VAULT_INFRA_RUN
```

You should now be able to login via OIDC on `https://vault.management.${management_infra_domain}` and see the configurations.
