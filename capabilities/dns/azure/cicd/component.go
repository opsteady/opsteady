package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilititesDNSAzure is an Azure implmentation for the external-dns controller
type CapabilitiesDNSAzure struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesDNSAzure component
func (c *CapabilitiesDNSAzure) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-azure", "kubernetes-azure-cluster", "capabilities-dns-azure")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.7.1", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}
