package component

import "github.com/opsteady/opsteady/cli/tasks"

// Validate validates the Terraform code
func (c *DefaultComponent) Validate() {

	executeOrder := c.DeterminOrderOfExecution()
	if len(c.OverrideValidateOrder) != 0 {
		executeOrder = c.OverrideValidateOrder
	}

	for _, folder := range executeOrder {
		switch folder {
		case c.Terraform:
			c.ValidateTerraform()
		}
	}
}

func (c *DefaultComponent) ValidateTerraform() {
	// Don't need to pass the storage name in, only a formatting check
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, "", c.GlobalConfig.CachePath, c.Logger)

	if err := terraform.FmtCheck(); err != nil {
		c.Logger.Fatal().Err(err).Msg("validation failed")
	}
}
