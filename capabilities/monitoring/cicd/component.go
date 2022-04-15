package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilititesMonitoring is an implementation for Prometheus
type CapabilititiesMonitoring struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilititesMonitoring component
func (c *CapabilititiesMonitoring) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.RequiresComponents("foundation-aws", "kubernetes-aws-cluster", "foundation-azure", "kubernetes-azure-cluster")
	c.SetVaultInfoToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"kube-prometheus-stack",
		"34.10.0", // renovate: datasource=helm registryUrl=https://https://prometheus-community.github.io/helm-charts depName=kube-prometheus-stack versioning=semver
	))
}
