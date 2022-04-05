package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilitiesCertificatesAWS is an AWS implementation for the certificates controller
type CapabilitiesCertificatesAWS struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesCertificatesAWS component
func (c *CapabilitiesCertificatesAWS) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.AddManagementCredentialsToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.8.0", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
