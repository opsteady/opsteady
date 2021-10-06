package component

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/tasks"
)

// Destroy runs the component destruction.
func (c *DefaultComponent) Destroy() {
	c.SetCloudCredentialsToEnv()
	c.SetPlatformInfoToComponentConfig()
	componentConfig := c.RetrieveComponentConfig()

	executeOrder := c.DetermineOrderOfExecution()
	if len(c.OverrideDestroyOrder) != 0 {
		executeOrder = c.OverrideDestroyOrder
	}

	for _, folder := range executeOrder {
		switch folder {
		case c.Terraform:
			c.PrepareTerraformBackend()
			c.DestroyTerraform(componentConfig)
		case c.Helm:
			c.LoginToAKSorEKS(componentConfig)
			c.DestroyHelm(componentConfig)
		}
	}
}

// DestroyTerraform destroyes resources created by Terrform
func (c *DefaultComponent) DestroyTerraform(values map[string]interface{}) {
	backendStorageName := values["management_bootstrap_terraform_state_account_name"].(string) // Always expecting this to be here
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)

	varsPath := fmt.Sprintf("%s/%s.tfvars.json", c.GlobalConfig.TmpFolder, c.ComponentName)
	c.WriteConfigToJSON(varsPath)

	if err := terraform.Destroy(varsPath); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}

// DestroyHelm removes Helm charts from Kubernetes
func (c *DefaultComponent) DestroyHelm(componentConfig map[string]interface{}) {
	helm := tasks.NewHelm(c.GlobalConfig.TmpFolder, c.Logger)
	for _, chart := range c.HelmCharts {
		if err := helm.Delete(chart.Release, chart.Namespace, c.DryRun); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not install Helm chart")
		}
	}
}
