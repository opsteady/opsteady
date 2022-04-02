package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilitiesLoadbalacing is an implementation for the Nginx controller
type CapabilitiesLoadbalacing struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesLoadbalacing component
func (c *CapabilitiesLoadbalacing) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"ingress-nginx",
		"4.0.19", // renovate: datasource=helm registryUrl=https://kubernetes.github.io/ingress-nginx depName=ingress-nginx versioning=semver
	))
}
