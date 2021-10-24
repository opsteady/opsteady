package component

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/tasks"
	"github.com/opsteady/opsteady/cli/templating"
)

// Deploy runs the component setup and deployment.
func (c *DefaultComponent) Deploy() {
	c.SetCloudCredentialsToEnv()
	c.SetPlatformInfoToComponentConfig()

	executeOrder := c.DetermineOrderOfExecution()
	if len(c.OverrideDeployOrder) != 0 {
		executeOrder = c.OverrideDeployOrder
	}

	for _, folder := range executeOrder {
		componentConfig := c.RetrieveComponentConfig()

		switch folder {
		case c.Terraform:
			c.PrepareTerraformBackend()
			c.DeployTerraform(componentConfig)
		case c.Helm:
			c.LoginToAKSorEKS(componentConfig)
			c.LoginToHelmRegistry()
			c.DeployHelm(componentConfig)
		case c.Kubectl:
			c.LoginToAKSorEKS(componentConfig)
			c.DeployKubectl(componentConfig)
		}
	}
}

// DeployTerraform uses Terrform code to deploy resources
func (c *DefaultComponent) DeployTerraform(componentConfig map[string]interface{}) {
	backendStorageName := componentConfig["management_bootstrap_terraform_state_account_name"].(string) // Always expecting this to be here
	terraform := tasks.NewTerraform(c.TerraformFolder(), c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)

	varsPath := fmt.Sprintf("%s/%s.tfvars.json", c.GlobalConfig.TmpFolder, c.ComponentName)
	c.WriteConfigToJSON(varsPath)

	if c.DryRun {
		c.Logger.Info().Msg("DryRun mode activated")
		if err := terraform.InitAndPlan(varsPath); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not plan")
		}
		return
	}

	if err := terraform.InitAndApply(varsPath); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}

	c.RetrieveComponentConfigWithoutCache() // We might have changed some outputs
}

// DeployHelm deploys Helm charts to Kubernetes
func (c *DefaultComponent) DeployHelm(componentConfig map[string]interface{}) {
	template := templating.NewTemplating(c.Logger)

	helm := tasks.NewHelm(c.Logger)
	for _, chart := range c.HelmCharts {
		if err := template.Render(c.HelmFolder(), c.HelmTmpFolder(chart.Release), componentConfig); err != nil {
			c.Logger.Fatal().Err(err).Str("chart", chart.Release).Msg("could not template Helm values files")
		}

		if err := helm.Upgrade(c.HelmTmpFolder(chart.Release), c.GlobalConfig.ManagementHelmRepository, chart.Release, chart.Namespace, chart.Version, c.DryRun); err != nil {
			c.Logger.Fatal().Err(err).Str("chart", chart.Release).Msg("could not install Helm chart")
		}
	}
}

// DeployKubectl deploys Kubernetes yaml files to Kubernetes
func (c *DefaultComponent) DeployKubectl(componentConfig map[string]interface{}) {
	template := templating.NewTemplating(c.Logger)

	if err := template.Render(c.KubectlFolder(), c.KubectlTmpFolder(), componentConfig); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not template Kubernetes manifest files")
	}

	kubectl := tasks.NewKubectl(c.Logger)
	if err := kubectl.Apply(c.KubectlTmpFolder(), c.DryRun); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply Kubernetes manifest files")
	}
}
