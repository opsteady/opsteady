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

**Note-1: Make sure to replace `dev-management` with whatever container registry name you will be using.**

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
# Run this from the root of the repository
docker run -it --rm -v $(pwd):/data dev-management.azurecr.io/cicd:1.0.0 /bin/bash
```

Before you start, comment the entire `backend "azurerm"` section in `management/bootstrap/terraform.tf`. This will keep the state local temporarily.

When bootstrapping we provide a set of defaults for the Terraform values. They are located in `management/defaults/bootstrap.default.tfvars.json`. Copy it to `management/defaults/bootstrap.tfvars.json` and adjust the values. This custom file will be ignored in the Git repository and is for your use only.

Actual steps to perform in the container locally, VSC remote container or on your local machine:

```bash
az login --use-device-code
az account set --subscription management
cd management/bootstrap
terraform init
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64

terraform apply -var-file=../defaults/bootstrap.tfvars.json
```

Uncomment the entire `backend "azurerm"` section in `management/bootstrap/terraform.tf` to upload the state to the remote backend.

```bash
terraform init -reconfigure -backend-config "storage_account_name=This name should match management_bootstrap_terraform_state_account_name"
```

`terraform.tfstate` will be empty and as it is uploaded to the remote storage, it can be safely deleted together with `terraform.tfstate.backup`.

## 03 Infra

When creating the initial management infra, we provide a set of defaults for the Terraform values. They are located in `management/defaults/infra.default.tfvars.json`. To deploy the management infrastructure successfully, copy the default file to a custom tfvars file called `management/defaults/infra.tfvars.json`.

In this file also update the `management_infra_key_vault_administrators` to include your AD user object ID. You can find this in the Azure AD. This makes sure that you can administer the key vault.
In this file also update the `management_infra_platform_admins` to include your AD user object ID. You can find this in the Azure AD. This makes sure that you can login to the Vault later on and manageme the entire platform.

```bash
cd management/infra
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64
terraform apply -var-file=../defaults/infra.tfvars.json

# Run the Terraform plan again to make sure that everything is applied.
# The virtual network rules in the Key Vault network ACL might not apply
# the first time for some unknown reason.
terraform apply -var-file=../defaults/infra.tfvars.json
```

At this point a DNS zone has been created for the Opsteady platform. From your DNS hosting provider you need to delegate a subzone to this domain.

```bash
terraform state show azurerm_dns_zone.public_root
```

The output of this command will show you the name servers (amongst other things) that you need to delegate to. The `management_infra_domain` variable in the infra defaults contains the subdomain that you need to delegate to. Create the NS records with your DNS hosting provider. It can take some time before the DNS resolving is active.

## 04 Vault Infrastructure

When creating the Vault infrastructure, we provide a set of defaults for the Terraform values. They are located in `management/defaults/vault-infra.default.tfvars.json`. To deploy the Vault infrastructure successfully, copy the default file to a custom tfvars file called `management/defaults/vault-infra.tfvars.json`. In this file update the `management_vault_infra_storage_account_name` to a unique name. This storage account will host the Vault CA certificate.

```bash
cd management/vault/infra
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64

# This command will give you some warnings about values for undeclared variables. This is expected and can be ignored.
terraform apply -compact-warnings -var-file=../../defaults/infra.tfvars.json -var-file=../../defaults/vault-infra.tfvars.json
```

Vault is now deployed to the cluster but all the instances are in a sealed state. We need initialise the cluster with the following steps:

```bash
# Get the admin credentials for the Kubernetes cluster
az aks get-credentials -g management -n ${management_infra_aks_name} --admin

# Initialize the Vault cluster
kubectl exec -n platform -it vault-0 -- vault operator init -ca-path=/vault/userconfig/vault-tls/ca.crt
```

**If the command succeeds you will see the recovery keys and the initial root token for Vault. Store this in a secure location and distribute the recovery keys to trusted parties.**

The certificate authority file for Vault can be downloaded from `https://${management_vault_infra_storage_account_name}.blob.core.windows.net/vault-ca/ca.pem`. With this file you should be able to connect securely to Vault on `https://vault.management.${management_infra_domain}`.

## 04 Vault Configuration

When creating the Vault configuration, we provide a set of defaults for the Terraform values. They are located in `management/defaults/vault-config.default.tfvars.json`. To deploy the Vault configuration successfully, copy the default file to a custom tfvars file called `management/defaults/vault-config.tfvars.json`. Add the management subscription ID as a value to the `management_vault_config_subscriptions` object.

```bash
cd management/vault/config
terraform init -backend-config storage_account_name="This name should match management_bootstrap_terraform_state_account_name"
terraform providers lock -platform=darwin_amd64 -platform=linux_amd64

# Before we can apply the Terraform code, we need to grab the Vault CA certificate and put it in a well-known location, so that the Vault provider can use it.
curl -o vault-ca.pem https://${management_vault_infra_storage_account_name}.blob.core.windows.net/vault-ca/ca.pem

# This command will give you some warnings about values for undeclared variables. This is expected and can be ignored.
terraform apply -compact-warnings -var-file=../../defaults/infra.tfvars.json -var-file=../../defaults/vault-config.tfvars.json -var vault_token=$VAULT_ROOT_TOKEN_FROM_VAULT_INFRA_RUN
```

You should now be able to login via OIDC on `https://vault.management.${management_infra_domain}` and see the configurations.

## 05 Revoke the Vault root token

The Vault root token should only be used in emergencies and never in regular Vault usage. The root token can always be regenerated from the recovery keys. To ensure maximum security you should revoke the root token with the following commands:

```bash
export VAULT_TOKEN=$ROOT_TOKEN_FROM_VAULT_INIT

vault token revoke -ca-cert=vault-ca.pem -address=https://vault.management.${management_infra_domain} -self
```

The initial manual bootstrap for the management environment is now complete, well done!

# Switching to the Opsteady CLI for management

**IMPORTANT: Please exit the container that you have been working in before executing the next steps**

## 01 Add Vault CA certificate to the CI/CD container image

The Vault CA is needed to securely connect to Vault from the CLI. The Vault CA is hosted at a well-known location that you've configured in the Vault infrastructure step. Build the Docker image again but this time add the a build argument to trigger the Vault CA download during build:

```bash
cd docker/cicd
docker build --build-arg ACR_NAME=dev-management --build-arg VAULT_CA_STORAGE_ACCOUNT=${management_vault_infra_storage_account_name} -t dev-management.azurecr.io/cicd:1.0.0 .
```

Note: If you want to work locally please add the certificate to your local machine, CLI and the browser. This works for Ubuntu:

```bash
curl -o vault-ca.pem https://$VAULT_CA_STORAGE_ACCOUNT.blob.core.windows.net/vault-ca/ca.pem
sudo openssl x509 -in vault-ca.pem -inform PEM -out /usr/local/share/ca-certificates/vault-ca.crt
sudo update-ca-certificates
```

## 02 Push the CI/CD container image to the registry

After a successful build you can push the image to the registry:

```bash
az login
az account set --subscription management
az acr login -n ${management_infra_acr_name}
docker push ${management_infra_acr_name}.azurecr.io/cicd:1.0.0 .
```

## 03 Seed Vault with management configuration

Before we can start using the CLI to manage our management environment we need to seed the Vault with our configuration data. For each of the components (management infra, vault infra and vault config) you have created a `$COMPONENT.fvars.json` file that contains the settings for Terraform. This information now needs to be stored in Vault. Please review the contents of the `tfvars.json` files in the `management/defaults` folder and make the adjustments you want.

```bash
# Make sure that we are in the root of the repository
cd ../../

# Start the CI/CD container (replace dev-management with your ACR name)
docker run -it --rm \
  -p 8250:8250 \
  -v $(pwd):/data \
  -v ${HOME}/.cache:/home/platform/.cache \
  -v ${HOME}/.cache/opsteady-go:/home/platform/go \
  dev-management.azurecr.io/cicd:1.0.0 /bin/bash

# Set the Vault address
export VAULT_ADDR=https://vault.management.${management_infra_domain}

# Log into the Vault with OIDC
# NOTE: make sure that you are part of the platform-admin group for this command to succeed. Also, Vault will not be able
# to open a browser from the container, so just click/copy the provided link and open it in the browser yourself.
vault login -method=oidc role=platform-admin listenaddress=0.0.0.0

# Store the component configurations
vault kv put config/v0/platform/management/management-bootstrap-default @management/defaults/bootstrap.tfvars.json
vault kv put config/v0/platform/management/management-infra-default @management/defaults/infra.tfvars.json
vault kv put config/v0/platform/management/management-vault-infra-default @management/defaults/vault-infra.tfvars.json
vault kv put config/v0/platform/management/management-vault-config-default @management/defaults/vault-config.tfvars.json
```

## 04 Run the CLI

We are now ready to manage all the management components with the CLI. First we need to configure the CLI by copying the `default-config.yaml` to `config.yaml` and entering the appropriate values for the Vault location and management environment.

```bash
cp default-config.yaml config.yaml

# Open the config.yaml and enter the correct values for your environment
```

Now we can run the CLI for all the components and check the Terraform plan. Depending on if you made any changes to the configurations, you will see some changes but mostly it should show none.

```bash
go run main.go deploy -c management-bootstrap --dry-run --azure-id management --cache
go run main.go deploy -c management-infra --dry-run --azure-id management --cache
go run main.go deploy -c management-vault-infra --dry-run --azure-id management --cache
go run main.go deploy -c management-vault-config --dry-run --azure-id management --cache
```
