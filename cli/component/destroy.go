package component

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/tasks"
	"github.com/opsteady/opsteady/cli/templating"
	"github.com/opsteady/opsteady/cli/utils"
)

// Destroy runs the component destruction.
// Run the destroy in the opposite direction than the deploy
// OverrideDestroyOrder is used as is if set
func (c *DefaultComponent) Destroy() {
	c.SetCloudCredentialsToEnv()
	c.SetPlatformInfoToComponentConfig()

	executeOrder := utils.ReverseStringArray(c.DetermineOrderOfExecution())
	if len(c.OverrideDeployOrder) != 0 {
		executeOrder = utils.ReverseStringArray(c.OverrideDeployOrder)
	}
	if len(c.OverrideDestroyOrder) != 0 {
		executeOrder = c.OverrideDestroyOrder
	}

	for _, folder := range executeOrder {
		componentConfig := c.RetrieveComponentConfig()

		switch folder {
		case c.Terraform:
			c.PrepareTerraformBackend()
			c.DestroyTerraform(componentConfig)
		case c.CRD:
			c.LoginToAKSorEKS(componentConfig)
			c.DestroyCRD(componentConfig)
		case c.KubeSetup:
			c.LoginToAKSorEKS(componentConfig)
			c.DestroyKubeSetup(componentConfig)
		case c.Helm:
			c.LoginToAKSorEKS(componentConfig)
			c.DestroyHelm(componentConfig)
		case c.KubePostSetup:
			c.LoginToAKSorEKS(componentConfig)
			c.DestroyKubePostSetup(componentConfig)
		}
	}
}

// DestroyTerraform destroyes resources created by Terrform
func (c *DefaultComponent) DestroyTerraform(values map[string]interface{}) {
	backendStorageName := values["management_bootstrap_terraform_state_account_name"].(string) // Always expecting this to be here
	terraform := tasks.NewTerraform(c.TerraformFolder(), c.TerraformBackendConfigPath, backendStorageName, c.GlobalConfig.CachePath, c.Logger)

	varsPath := fmt.Sprintf("%s/%s.tfvars.json", c.GlobalConfig.TmpFolder, c.ComponentName)
	c.WriteConfigToJSON(varsPath)

	if err := terraform.Destroy(varsPath); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not apply")
	}
}

// DestroyHelm removes Helm charts from Kubernetes
func (c *DefaultComponent) DestroyHelm(componentConfig map[string]interface{}) {
	helm := tasks.NewHelm(c.Logger)
	for _, chart := range c.HelmCharts {
		if err := helm.Delete(c.HelmTmpFolder(chart.Release), chart.Release, chart.Namespace, c.DryRun); err != nil {
			c.Logger.Fatal().Err(err).Msg("could not install Helm chart")
		}
	}
}

// DestroyKubeSetup destroy Kubernetes yaml files to Kubernetes
func (c *DefaultComponent) DestroyKubeSetup(componentConfig map[string]interface{}) {
	template := templating.NewTemplating(c.Logger)

	if err := template.Render(c.KubeSetupFolder(), c.KubeSetupTmpFolder(), componentConfig); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not template Kubernetes manifest files")
	}

	kubectl := tasks.NewKubectl(c.Logger)
	if err := kubectl.Delete(c.KubeSetupTmpFolder(), c.DryRun); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not delete Kubernetes manifest files")
	}
}

// DestroyKubePostSetup destroy Kubernetes yaml files to Kubernetes
func (c *DefaultComponent) DestroyKubePostSetup(componentConfig map[string]interface{}) {
	template := templating.NewTemplating(c.Logger)

	if err := template.Render(c.KubePostSetupFolder(), c.KubePostSetupTmpFolder(), componentConfig); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not template Kubernetes manifest files")
	}

	kubectl := tasks.NewKubectl(c.Logger)
	if err := kubectl.Delete(c.KubePostSetupTmpFolder(), c.DryRun); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not delete Kubernetes manifest files")
	}
}

// DestroyCRD destroy CRD yaml files to Kubernetes
func (c *DefaultComponent) DestroyCRD(componentConfig map[string]interface{}) {
	kubectl := tasks.NewKubectl(c.Logger)
	if err := kubectl.Delete(c.CRDFolder(), c.DryRun); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not delete CRD manifest files")
	}
}
