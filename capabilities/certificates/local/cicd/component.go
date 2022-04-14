package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilitiesCertificatesLocal is a local implementation for the certificates controller
type CapabilitiesCertificatesLocal struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesCertificatesLocal component
func (c *CapabilitiesCertificatesLocal) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-local", "kubernetes-local-cluster")
	c.DefaultComponent.AddAzureADCredentialsToComponentConfig()
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.8.0", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
