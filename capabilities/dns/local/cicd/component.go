package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationLocal "github.com/opsteady/opsteady/foundation/local/cicd"
	kubernetesLocalCluster "github.com/opsteady/opsteady/kubernetes/local/cluster/cicd"
)

// CapabilitiesDNSLocal is a local implementation for the external-dns controller
type CapabilitiesDNSLocal struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesDNSLocal{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "dns"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetLocal)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesDNSLocal before running
func (c *CapabilitiesDNSLocal) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationLocal.Instance.GetMetadata(), kubernetesLocalCluster.Instance.GetMetadata())
	c.SetVaultInfoToComponentConfig()
	c.AddAzureADCredentialsToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.13.0", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}
