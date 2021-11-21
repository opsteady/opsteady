package component

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opsteady/opsteady/cli/configuration"
	"github.com/opsteady/opsteady/cli/credentials"
	"github.com/opsteady/opsteady/cli/tasks"
	"github.com/opsteady/opsteady/cli/utils"
	"github.com/opsteady/opsteady/cli/vault"
	"github.com/rs/zerolog"
)

const (
	management = "management"
)

// Component is the interface all components have to implement.
type Component interface {
	Validate()
	Build()
	Deploy()
	Destroy()
	Test()
	Clean()
	Publish()
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
	PlatformVersion       string           // Version of the platform (used as a folder in Vault)
	HelmCharts            []*HelmChart     // We expect all charts to be from management ACR
	DockerBuildInfo       *DockerBuildInfo // We expect all docker images to be saved to ACR
	// Names of the folders a component uses which will determin what will be executed, order can be adjusted
	Terraform     string
	CRD           string
	KubeSetup     string
	Helm          string
	KubePostSetup string
	Docker        string
	// Override order
	OverrideDeployOrder   []string
	OverrideDestroyOrder  []string
	OverrideValidateOrder []string
}

// Test runs the component tests.
func (c *DefaultComponent) Test() {
	c.Logger.Warn().Msg("Test not implemented")
}

// RequiresComponents sets other components this component requires on.
func (c *DefaultComponent) RequiresComponents(dependencies ...string) {
	c.ComponentDependencies = dependencies
}

// RequiresComponents sets other components this component requires on.
func (c *DefaultComponent) SetDockerBuildInfo(name, version string, buildArgs map[string]string) {
	c.DockerBuildInfo = &DockerBuildInfo{
		Name:      name,
		Version:   version,
		BuildArgs: buildArgs,
	}
}

// SetCloudCredentialsToEnv gets the AWS or Azure credentials and
// sets them to env so they can be used by the CLI further down the process.
func (c *DefaultComponent) SetCloudCredentialsToEnv() {
	if c.AwsID != "" {
		c.setAwsCloudCredentialsToEnv()
	}

	if c.AzureID != "" {
		c.setAzureCloudCredentialsToEnv()
	}
}

func (c *DefaultComponent) setAwsCloudCredentialsToEnv() {
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

func (c *DefaultComponent) setAzureCloudCredentialsToEnv() {
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

// AzureIDorAwsID returns AzureID if both AWS and Azure ID are set
func (c *DefaultComponent) AzureIDorAwsID() string {
	if c.AzureID == "" {
		return c.AwsID
	}
	return c.AzureID
}

func (c *DefaultComponent) ComponentNameAndAllTheDependencies() []string {
	return utils.UniqueNonEmptyElementsOf(append([]string{c.ComponentName}, append(c.DefaultDependencies, c.ComponentDependencies...)...))
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

// RetrieveComponentConfig returns component config
func (c *DefaultComponent) RetrieveComponentConfig() map[string]interface{} {
	values, err := c.ComponentConfig.RetrieveConfig(c.PlatformVersion, c.AzureIDorAwsID(), c.ComponentNameAndAllTheDependencies())
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not retrieve component configuration")
	}
	return values
}

// AddAzureADCredentialsToComponentConfig adds Azure AD credentials to component config to be used elsewhere
func (c *DefaultComponent) AddAzureADCredentialsToComponentConfig() {
	azureAD, err := c.Credentials.AzureAD()
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not get credentials for AzureAD")
	}

	c.ComponentConfig.GeneralAddOrOverride("azuread_client_id", azureAD["client_id"].(string))
	c.ComponentConfig.GeneralAddOrOverride("azuread_client_secret", azureAD["client_secret"].(string))
	c.ComponentConfig.GeneralAddOrOverride("azuread_tenant_id", c.GlobalConfig.TenantID)
}

// AddManagementCredentialsToComponentConfig adds management credentials to component config to be used elsewhere
func (c *DefaultComponent) AddManagementCredentialsToComponentConfig() {
	mgmtSubscriptionCreds, err := c.Credentials.Azure(management)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not get credentials for management subscription")
	}

	c.ComponentConfig.GeneralAddOrOverride("management_client_id", mgmtSubscriptionCreds["client_id"].(string))
	c.ComponentConfig.GeneralAddOrOverride("management_client_secret", mgmtSubscriptionCreds["client_secret"].(string))
	c.ComponentConfig.GeneralAddOrOverride("management_subscription_id", c.GlobalConfig.ManagementSubscriptionID)
	c.ComponentConfig.GeneralAddOrOverride("tenant_id", c.GlobalConfig.TenantID)
}

// LoginToAKSorEKS logs in to AKS or EKS
func (c *DefaultComponent) LoginToAKSorEKS(componentConfig map[string]interface{}) {
	if c.AwsID != "" {
		aws := tasks.NewAws(c.GlobalConfig.TmpFolder, c.Logger)

		if err := aws.LoginToEKS(componentConfig["foundation_aws_region"].(string), componentConfig["kubernetes_aws_cluster_name"].(string)); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not login to EKS")
		}
	}

	if c.AzureID != "" {
		c.Logger.Info().Msg("Preparing AKS environment...")
		AKSCreds, err := c.Credentials.AKS(c.AzureID)

		if err != nil {
			c.Logger.Fatal().Err(err).Msg("could not get credentials to prepare AKS")
		}

		azTask := tasks.NewAz(c.GlobalConfig.TmpFolder, c.Logger)

		if err := azTask.LoginToAzure(AKSCreds["client_id"].(string), AKSCreds["client_secret"].(string), c.GlobalConfig.TenantID); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not login to Azure")
		}
		clusterName := componentConfig["kubernetes_azure_cluster_name"].(string)
		clusterResourceGroup := fmt.Sprintf("kubernetes-%s", clusterName)
		// Management cluster is different therefore we override this stuff here
		if clusterName == management {
			clusterResourceGroup = management
		}

		if err := azTask.LoginToAKS(clusterName, clusterResourceGroup); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not login to ASK via az")
		}
	}
}

// WriteConfigToJSON marshalls the component configuration to JSON format and
// writes it to a file indicated by the path parameter.
func (c *DefaultComponent) WriteConfigToJSON(path string) {
	config, err := json.Marshal(c.RetrieveComponentConfig())

	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not marshall the component configuration to JSON")
	}

	err = os.WriteFile(path, config, 0644)

	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not write the component configuration JSON to a file")
	}
}
