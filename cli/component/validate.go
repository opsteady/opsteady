package component

import "github.com/opsteady/opsteady/cli/tasks"

// Validate validates the Terraform code
func (c *DefaultComponent) Validate() {
	executeOrder := c.DetermineOrderOfExecution()
	if len(c.OverrideValidateOrder) != 0 {
		executeOrder = c.OverrideValidateOrder
	}

	for _, folder := range executeOrder {
		switch folder {
		case c.Terraform:
			c.ValidateTerraform()
		case c.Docker:
			c.ValidateDocker()
		}
	}
}

func (c *DefaultComponent) ValidateTerraform() {
	// Don't need to pass the storage name in, only a formatting check
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, "", c.GlobalConfig.CachePath, c.Logger)

	if err := terraform.FmtCheck(); err != nil {
		c.Logger.Fatal().Err(err).Msg("terraform validation failed")
	}
}

func (c *DefaultComponent) ValidateDocker() {
	docker := tasks.NewDocker(c.Logger)

	if err := docker.Validate(c.DockerFolder()); err != nil {
		c.Logger.Fatal().Err(err).Msg("docker validation failed")
	}
}
