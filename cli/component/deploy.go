package component

import (
	"github.com/opsteady/opsteady/cli/tasks"
)

// Deploy runs the component setup and deployment.
func (c *DefaultComponent) Deploy() {
	c.SetCloudCredentialsToEnv()
	c.SetVaultInfoToComponentConfig()
	c.SetPlatformInfoToComponentConfig()
	componentConfig := c.RetrieveComponentConfig()

	executeOrder := c.DeterminOrderOfExecution()
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
func (c *DefaultComponent) DeployTerraform(componentConfig map[string]string) {
	backendStorageName := componentConfig["management_bootstrap_terraform_state_account_name"] // Always expecting this to be here
	terraform := tasks.NewTerraform(c.ComponentFolder, c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)
	if c.DryRun {
		c.Logger.Info().Msg("DryRun........")
		if err := terraform.InitAndPlan(componentConfig); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not plan")
		}
		return
	}

	if err := terraform.InitAndApply(componentConfig); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}

// DeployHelm deploys Helm charts to Kubernetes
func (c *DefaultComponent) DeployHelm(componentConfig map[string]string) {
	// TODO: Add templating to values.yaml here before executing and pass the templated values.yaml file to the Helm install
	helm := tasks.NewHelm(c.GlobalConfig.TmpFolder, c.Logger)
	for _, chart := range c.HelmCharts {
		// TODO: this is just an example, not tested yet
		if err := helm.Upgrade(c.GlobalConfig.ManagementHelmRepository, chart.Release, chart.Namespace, chart.Version, c.DryRun); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not install Helm chart")
		}
	}
}
