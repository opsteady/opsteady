package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilititesLogging is an implementation for Prometheus
type CapabilititiesLogging struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilititesLogging component
func (c *CapabilititiesLogging) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.RequiresComponents("foundation-aws", "kubernetes-aws-cluster", "foundation-azure", "kubernetes-azure-cluster")
	c.UseHelm(
		component.NewHelmChart(
			"loki",
			"2.11.0", // renovate: datasource=helm registryUrl=https://grafana.github.io/helm-charts depName=loki versioning=semver
		),
		component.NewHelmChart(
			"promtail",
			"4.2.0", // renovate: datasource=helm registryUrl=https://grafana.github.io/helm-charts depName=promtail versioning=semver
		),
	)
}
