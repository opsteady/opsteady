package cicd

import (
	"github.com/opsteady/opsteady/cli/component"

	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// CapabilitiesLoadbalacing is an implementation for the Nginx controller
type CapabilitiesLoadbalacing struct {
	component.DefaultComponent
	Metadata *component.Metadata
}

var Instance = &CapabilitiesLoadbalacing{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "loadbalacing"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetLocal, component.TargetAws, component.TargetAzure)
	m.AddGroupDependency(component.KubernetesAddons) // Local misses Addons but we can for now calculate the dependency
	Instance.Metadata = &m
}

// Configure configures CapabilitiesLoadbalacing before running
func (c *CapabilitiesLoadbalacing) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	c.UseHelm(component.NewHelmChart(
		"ingress-nginx",
		"4.0.17", // renovate: datasource=helm registryUrl=https://kubernetes.github.io/ingress-nginx depName=ingress-nginx versioning=semver
	))
}
