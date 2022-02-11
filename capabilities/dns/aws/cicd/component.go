package cicd

import "github.com/opsteady/opsteady/cli/component"

// CapabilititesDNSAWS is an AWS implmentation for the external-dns controller
type CapabilitiesDNSAWS struct {
	component.DefaultComponent
}

// Initialize creates a new CapabilitiesDNSAWS component
func (c *CapabilitiesDNSAWS) Initialize(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	c.DefaultComponent.SetVaultInfoToComponentConfig()
	c.DefaultComponent.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.7.1", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}

func (k *CapabilitiesDNSAWS) Info() component.ComponentDepInfo {
	return component.ComponentDepInfo{
		Description:    "Creates EKS",
		Group:          "Kubernetes Cluster",
		DependsOn:      []string{""},
		DependsOnGroup: "",
	}
}
