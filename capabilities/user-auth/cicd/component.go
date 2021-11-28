package cicd

import "github.com/opsteady/opsteady/cli/component"

// UserAuth is an implementation for the user authentication
type UserAuth struct {
	component.DefaultComponent
}

// Initialize creates a new UserAuth component
func (c *UserAuth) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents(
		"foundation-azure",
		"foundation-aws",
		"kubernetes-azure-cluster",
		"kubernetes-aws-cluster")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.AddManagementCredentialsToComponentConfig()
	c.DefaultComponent.AddAzureADCredentialsToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"dex",
		"0.6.3", // renovate: datasource=helm registryUrl=https://charts.dexidp.io depName=dex versioning=semver
	))
}
