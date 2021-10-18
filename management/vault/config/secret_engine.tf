# Mount the kv (v2) secret backend at /config. We use this backend to store all information
# related to the platform accounts and subscriptions, including management.
resource "vault_mount" "kv" {
  path = "config"
  type = "kv-v2"
}

# AzureAD secrets engine

module "vault_azuread_access" {
  source = "../../../internal/modules/service-principal"
  name   = "vault-azuread-access"

  required_resource_access = [
    {
      resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

      resource_access = [
        {
          id   = azuread_service_principal.msgraph.app_role_ids["Application.ReadWrite.All"]
          type = "Role"
        },
        {
          id   = azuread_service_principal.msgraph.app_role_ids["Directory.Read.All"]
          type = "Role"
        }
      ]
    },
    {
      resource_app_id = data.azuread_application_published_app_ids.well_known.result.AzureActiveDirectoryGraph

      resource_access = [
        {
          id   = azuread_service_principal.aadgraph.app_role_ids["Application.ReadWrite.All"]
          type = "Role"
        },
        {
          id   = azuread_service_principal.aadgraph.app_role_ids["Directory.Read.All"]
          type = "Role"
        }
      ]
    }
  ]
}

resource "vault_azure_secret_backend" "azuread_secrets" {
  subscription_id = data.azurerm_client_config.current.subscription_id
  tenant_id       = data.azurerm_client_config.current.tenant_id
  client_id       = module.vault_azuread_access.azuread_service_principal_application_id
  client_secret   = module.vault_azuread_access.azuread_service_principal_password
  path            = "azuread"
}

resource "vault_azure_secret_backend_role" "azuread_role" {
  backend               = vault_azure_secret_backend.azuread_secrets.path
  role                  = "management"
  application_object_id = module.vault_azuread_access.azuread_application_object_id
  ttl                   = 3600
  max_ttl               = 14400
}

# Assign the vault-azuread-secrets service principal the owner role on the management subscription, so that
# it can generate dynamic secrets for this subscription.
resource "azurerm_role_assignment" "azuread_secrets" {
  scope                = "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  role_definition_name = "Owner"
  principal_id         = module.vault_azuread_access.azuread_service_principal_object_id
}

# Azure Subscriptions secrets engine

# The vault-azure-access service principal is used to generate new service principals and
# assign ownership to Azure subscriptions so that they can be managed.
module "vault_azure_access" {
  source = "../../../internal/modules/service-principal"
  name   = "vault-azure-access"

  required_resource_access = [
    {
      resource_app_id = data.azuread_application_published_app_ids.well_known.result.AzureActiveDirectoryGraph

      resource_access = [
        {
          id   = azuread_service_principal.aadgraph.app_role_ids["Application.ReadWrite.All"]
          type = "Role"
        }
      ]
    }
  ]
}

# Creation of service principals takes time to propagate within Azure AD.
resource "time_sleep" "wait_for_service_principal_propagation" {
  depends_on = [module.vault_azure_access, module.vault_azuread_access]

  create_duration = "20s"
}

# This resource is expected to only run once. It grants admin consent on the Vault azure secrets application registration
# to create new applications. This command can only succeed if the identity running Terraform has admin credentials on the
# AD tenant. This is normally done the first time the management environment is locally bootstrapped by an admin user.
resource "null_resource" "grant_admin_consent" {

  # Grant Admin consent via the AZ CLI.
  provisioner "local-exec" {
    command = "az ad app permission admin-consent --id ${module.vault_azure_access.azuread_application_application_id}"
  }

  # Grant Admin consent via the AZ CLI.
  provisioner "local-exec" {
    command = "az ad app permission admin-consent --id ${module.vault_azuread_access.azuread_application_application_id}"
  }

  depends_on = [time_sleep.wait_for_service_principal_propagation]
}

# Assign the vault-azure-secrets service principlal the owner role on the management subscription, so that
# it can generate dynamic secrets for this subscription.
resource "azurerm_role_assignment" "azure_secrets" {
  scope                = "/subscriptions/${data.azurerm_client_config.current.subscription_id}"
  role_definition_name = "Owner"
  principal_id         = module.vault_azure_access.azuread_service_principal_object_id
}

# Mount the azure secret backend at /azure.
resource "vault_azure_secret_backend" "azure" {
  subscription_id = data.azurerm_client_config.current.subscription_id
  tenant_id       = data.azurerm_client_config.current.tenant_id
  client_id       = module.vault_azure_access.azuread_service_principal_application_id
  client_secret   = module.vault_azure_access.azuread_service_principal_password
  path            = "azure"
}

resource "vault_azure_secret_backend_role" "azure_owner_role" {
  for_each = var.management_vault_config_subscriptions

  backend = vault_azure_secret_backend.azure.path
  role    = each.key
  ttl     = 3600
  max_ttl = 14400

  azure_roles {
    role_name = "Owner"
    scope     = "/subscriptions/${each.value}"
  }

  # This role assignement ensures that the service principal can manage
  # the data that is in a key vault.
  azure_roles {
    role_name = "Key Vault Administrator"
    scope     = "/subscriptions/${each.value}"
  }

  depends_on = [azurerm_role_assignment.azure_secrets]
}

resource "vault_azure_secret_backend_role" "azure_kubernetes_role" {
  for_each = var.management_vault_config_subscriptions

  backend = vault_azure_secret_backend.azure.path
  role    = "${each.key}-k8s"
  ttl     = 3600
  max_ttl = 14400

  azure_roles {
    role_name = "Contributor"
    scope     = "/subscriptions/${each.value}"
  }

  azure_roles {
    role_name = "Azure Kubernetes Service Cluster Admin Role"
    scope     = "/subscriptions/${each.value}"
  }

  azure_roles {
    role_name = "Azure Kubernetes Service Cluster User Role"
    scope     = "/subscriptions/${each.value}"
  }

  depends_on = [azurerm_role_assignment.azure_secrets]
}

resource "vault_aws_secret_backend" "aws" {
  for_each = { for account in var.management_vault_config_accounts : account.name => account }

  access_key = each.value.access_key
  secret_key = each.value.secret_key
  region     = var.management_vault_config_aws_region
  path       = "aws/${each.key}"

  default_lease_ttl_seconds = 3600
  max_lease_ttl_seconds     = 14400
}

resource "vault_aws_secret_backend_role" "vault_aws_access_role" {
  for_each = { for account in var.management_vault_config_accounts : account.name => account }

  backend         = vault_aws_secret_backend.aws[each.key].path
  name            = "vault-aws-access"
  credential_type = "assumed_role"
  role_arns       = ["arn:aws:iam::${each.value.id}:role/vault-aws-access"] # This role will be configured with AWS organizations
  default_sts_ttl = 3600
  max_sts_ttl     = 14400
}
