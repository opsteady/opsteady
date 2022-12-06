package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilititesDNSLocal is a local implementation for the external-dns controller
type CapabilitiesDNSLocal struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesDNSLocal component
func (c *CapabilitiesDNSLocal) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-local", "kubernetes-local-cluster")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.AddAzureADCredentialsToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.12.0", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}
