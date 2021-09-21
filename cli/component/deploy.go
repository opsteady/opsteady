package component

import (
	"github.com/opsteady/opsteady/cli/tasks"
)

// Deploy Component
func (c *DefaultComponent) Deploy() {
	c.SetCloudCredentialsToEnv()
	c.PrepareTerraformBackend()
	c.SetVaultInfoToComponentConfig()
	c.SetPlatformInfoToComponentConfig()

	values, err := c.ComponentConfig.RetrieveConfig(c.PlatformVersion, c.AzureIDorAwsID(), c.ComponentNameAndAllTheDependencies())
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not deploy")
	}

	backendStorageName := values["management_bootstrap_terraform_state_account_name"] // Always expecting this to be here
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)
	if c.DryRun {
		c.Logger.Info().Msg("DryRun........")
		if err := terraform.InitAndPlan(values); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not plan")
		}
		return
	}

	if err := terraform.InitAndApply(values); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}
