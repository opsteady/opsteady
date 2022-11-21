package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilitiesCertificatesAzure is an Azure implementation for the certificates controller
type CapabilitiesCertificatesAzure struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesCertificatesAzure component
func (c *CapabilitiesCertificatesAzure) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-azure", "kubernetes-azure-cluster")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.10.1", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
