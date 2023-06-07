package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// CapabilitiesDNSAzure is an Azure implementation for the external-dns controller
type CapabilitiesDNSAzure struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesDNSAzure{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "dns"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetAzure)
	m.AddGroupDependency(component.KubernetesAddons)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesDNSAzure before running
func (c *CapabilitiesDNSAzure) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationAzure.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata())
	c.SetVaultInfoToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.13.0", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}
