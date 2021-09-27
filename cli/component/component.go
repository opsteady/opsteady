package component

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/opsteady/opsteady/cli/configuration"
	"github.com/opsteady/opsteady/cli/credentials"
	"github.com/opsteady/opsteady/cli/utils"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/rs/zerolog"
)

// Component is the interface all components have to implement.
type Component interface {
	Validate()
	Build()
	Deploy()
	Destroy()
	Test()
	Clean()
	Release()
}

// Initialize is an interface that ensures component initialization.
type Initialize interface {
	Initialize(DefaultComponent)
}

// DefaultComponent implements the Component interface in a general way.
// This implementation should suit all the components but if not, they
// can adjust it the way they want or even override the whole function.
// The implemented functions stop (Fatal) if an error is detected.
type DefaultComponent struct {
	// Component dependencies
	Vault                      vault.Vault
	Credentials                credentials.Credentials
	ComponentConfig            ComponentConfig
	GlobalConfig               *configuration.GlobalConfig
	Logger                     *zerolog.Logger
	TerraformBackendConfigPath string
	// Component configuration
	DefaultDependencies   []string // Dependencies that all components must have (used for fetching data from Vault)
	ComponentDependencies []string // Dependencies component has (used for fetching data from Vault)
	ComponentName         string
	ComponentFolder       string
	DryRun                bool
	AwsID                 string
	AzureID               string
	PlatformVersion       string // Version of the platform (used as a folder in Vault)
}

// RequiresComponents sets other components this component requires on.
func (c *DefaultComponent) RequiresComponents(dependencies ...string) {
	c.ComponentDependencies = dependencies
}

// SetCloudCredentialsToEnv gets the AWS or Azure credentials and
// sets them to env so they can be used by the CLI further down the process.
func (c *DefaultComponent) SetCloudCredentialsToEnv() {
	if c.AwsID != "" {
		awsAccountCreds, err := c.Credentials.AWS(c.AwsID, "60m")
		if err != nil {
			c.Logger.Fatal().Err(err).Str("awsID", c.AwsID).Msg("Could not get credentials")
		}

		if err := os.Setenv("AWS_ACCESS_KEY_ID", awsAccountCreds["access_key"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env AWS_ACCESS_KEY_ID")
		}

		if err := os.Setenv("AWS_SECRET_ACCESS_KEY", awsAccountCreds["secret_key"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env AWS_SECRET_ACCESS_KEY")
		}

		if err := os.Setenv("AWS_SESSION_TOKEN", awsAccountCreds["security_token"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env AWS_SESSION_TOKEN")
		}
	}
	if c.AzureID != "" {
		azureSubscriptionCreds, err := c.Credentials.Azure(c.AzureID)
		if err != nil {
			c.Logger.Fatal().Err(err).Str("azureID", c.AzureID).Msg("Could not get credentials")
		}

		if err := os.Setenv("ARM_CLIENT_ID", azureSubscriptionCreds["client_id"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env ARM_CLIENT_ID")
		}

		if err := os.Setenv("ARM_CLIENT_SECRET", azureSubscriptionCreds["client_secret"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env ARM_CLIENT_SECRET")
		}

		if err := os.Setenv("ARM_TENANT_ID", c.GlobalConfig.TenantID); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env ARM_TENANT_ID")
		}

		if err := os.Setenv("ARM_SUBSCRIPTION_ID", c.GlobalConfig.ManagementSubscriptionID); err != nil {
			c.Logger.Fatal().Err(err).Msg("Could not set env ARM_SUBSCRIPTION_ID")
		}
	}
}

// PrepareTerraformBackend prepares the terraform to use remote storage
// in the management subscription.
func (c *DefaultComponent) PrepareTerraformBackend() {
	c.Logger.Info().Msg("Preparing Terraform environment...")

	mgmtCreds, err := c.Credentials.Azure("management")
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not get management credentials to prepare terraform")
	}

	tfBackendCreds := fmt.Sprintf(`
	subscription_id = "%s"
	tenant_id = "%s"
	client_id = "%s"
	client_secret = "%s"`,
		c.GlobalConfig.ManagementSubscriptionID,
		c.GlobalConfig.TenantID,
		mgmtCreds["client_id"].(string),
		mgmtCreds["client_secret"].(string))

	// Try to determine which blob key to use for Terraform state.
	var blobKey string
	if c.AwsID != "" && c.AzureID == "" {
		blobKey = fmt.Sprintf("%s/%s/%s.tfstate", "aws", c.AwsID, c.ComponentName)
	} else if c.AwsID == "" && c.AzureID != "" {
		blobKey = fmt.Sprintf("%s/%s/%s.tfstate", "azure", c.AzureID, c.ComponentName)
	} else if c.AwsID == "" && c.AzureID == "" {
		c.Logger.Fatal().Msg("Please specify a target AWS/Azure ID")
	} else if c.AwsID != "" && c.AzureID != "" {
		c.Logger.Info().Msg("You specified both an Azure and AWS ID, using the backend blob key in the Terraform provider")
	}

	if blobKey != "" {
		c.Logger.Info().Str("backend", blobKey).Msg("Using backend blob key")
		tfBackendCreds = fmt.Sprintf("key = \"%s\"\n%s", blobKey, tfBackendCreds)
	}

	if err := ioutil.WriteFile(c.TerraformBackendConfigPath, []byte(tfBackendCreds), 0644); err != nil {
		c.Logger.Fatal().Err(err).Str("path", c.TerraformBackendConfigPath).Msg("could not write the backend config file")
	}
}

// AzureIDorAwsID returns AzureID if both AWS and Azure ID are set.
func (c *DefaultComponent) AzureIDorAwsID() string {
	if c.AzureID == "" {
		return c.AwsID
	}
	return c.AzureID
}

func (c *DefaultComponent) ComponentNameAndAllTheDependencies() []string {
	return utils.UniqueNonEmptyElementsOf(append([]string{c.ComponentName}, append(c.DefaultDependencies, c.ComponentDependencies...)...))
}

// Build generates an artifact for the component that can be released.
func (c *DefaultComponent) Build() {
	c.Logger.Warn().Msg("Build not implemented")
}

// Test runs the component tests.
func (c *DefaultComponent) Test() {
	c.Logger.Warn().Msg("Test not implemented")
}

// Release publishes the artifact that was generated by the build.
func (c *DefaultComponent) Release() {
	c.Logger.Warn().Msg("Release not implemented")
}

// SetVaultInfoToComponentConfig sets Vault address and token to ComponentConfig
// so that other steps can use it.
func (c *DefaultComponent) SetVaultInfoToComponentConfig() {
	c.ComponentConfig.GeneralAddOrOverride("vault_address", c.Vault.GetAddress())
	c.ComponentConfig.GeneralAddOrOverride("vault_token", c.Vault.GetToken())
}

// SetPlatformInfoToComponentConfig sets platform version, environment and component name
// to ComponentConfig so that other steps can use it.
func (c *DefaultComponent) SetPlatformInfoToComponentConfig() {
	c.ComponentConfig.GeneralAddOrOverride("platform_version", c.PlatformVersion)
	c.ComponentConfig.GeneralAddOrOverride("platform_environment_name", c.AzureIDorAwsID())
	c.ComponentConfig.GeneralAddOrOverride("platform_component_name", c.ComponentName)
}
