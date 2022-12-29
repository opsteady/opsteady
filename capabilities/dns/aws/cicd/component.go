package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// CapabilitiesDNSAWS is an AWS implementation for the external-dns controller
type CapabilitiesDNSAWS struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesDNSAWS{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "dns"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.KubernetesAddons)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesDNSAWS before running
func (c *CapabilitiesDNSAWS) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	c.SetVaultInfoToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"external-dns",
		"1.12.0", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/external-dns depName=external-dns versioning=semver
	))
}
