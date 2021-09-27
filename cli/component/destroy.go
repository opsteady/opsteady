package component

import "github.com/opsteady/opsteady/cli/tasks"

// Destroy runs the component destruction.
func (c *DefaultComponent) Destroy() {
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

	if err := terraform.Destroy(values); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}
