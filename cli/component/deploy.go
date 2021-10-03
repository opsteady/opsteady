package component

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/opsteady/opsteady/cli/tasks"
)

// Deploy runs the component setup and deployment.
func (c *DefaultComponent) Deploy() {
	c.SetCloudCredentialsToEnv()
	c.SetPlatformInfoToComponentConfig()
	componentConfig := c.RetrieveComponentConfig()

	executeOrder := c.DetermineOrderOfExecution()
	if len(c.OverrideDeployOrder) != 0 {
		executeOrder = c.OverrideDeployOrder
	}

	for _, folder := range executeOrder {
		switch folder {
		case c.Terraform:
			c.PrepareTerraformBackend()
			c.DeployTerraform(componentConfig)
		case c.Helm:
			c.LoginToAKSorEKS(componentConfig)
			c.LoginToHelmRegistry()
			c.DeployHelm(componentConfig)
		}
	}
}

// DeployTerraform uses Terrform code to deploy resources
func (c *DefaultComponent) DeployTerraform(componentConfig map[string]interface{}) {
	backendStorageName := componentConfig["management_bootstrap_terraform_state_account_name"].(string) // Always expecting this to be here
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)

	// Marshall the component configuration to a JSON tfvars file
	tfvars, err := json.Marshal(componentConfig)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("Failed to marshall the component config to tfvars JSON")
	}

	varsPath := fmt.Sprintf("/tmp/%s.tfvars.json", c.ComponentName)

	err = os.WriteFile(varsPath, tfvars, 0644)
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("Failed to create tfvars JSON file")
	}

	if c.DryRun {
		c.Logger.Info().Msg("DryRun........")
		if err := terraform.InitAndPlan(varsPath); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not plan")
		}
		return
	}

	if err := terraform.InitAndApply(varsPath); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}

// DeployHelm deploys Helm charts to Kubernetes
func (c *DefaultComponent) DeployHelm(componentConfig map[string]interface{}) {
	// TODO: Add templating to values.yaml here before executing and pass the templated values.yaml file to the Helm install
	helm := tasks.NewHelm(c.GlobalConfig.TmpFolder, c.Logger)
	for _, chart := range c.HelmCharts {
		// TODO: this is just an example, not tested yet
		if err := helm.Upgrade(c.GlobalConfig.ManagementHelmRepository, chart.Release, chart.Namespace, chart.Version, c.DryRun); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not install Helm chart")
		}
	}
}
