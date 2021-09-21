package component

import "github.com/opsteady/opsteady/cli/tasks"

// Validate validates the Terraform code
func (c *DefaultComponent) Validate() {
	// Don't need to pass the storage name in, only a formatting check
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, "", c.GlobalConfig.CachePath, c.Logger)

	if err := terraform.FmtCheck(); err != nil {
		c.Logger.Fatal().Err(err).Msg("validation failed")
	}
}
