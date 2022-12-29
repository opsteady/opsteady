package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// CapabilitiesLoadBalancing is an implementation for the Nginx controller
type CapabilitiesLoadBalancing struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesLoadBalancing{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "loadBalancing"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetLocal, component.TargetAws, component.TargetAzure)
	m.AddGroupDependency(component.KubernetesAddons) // Local misses Addons but we can for now calculate the dependency
	Instance.Metadata = &m
}

// Configure configures CapabilitiesLoadBalancing before running
func (c *CapabilitiesLoadBalancing) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	c.UseHelm(component.NewHelmChart(
		"ingress-nginx",
		"4.4.0", // renovate: datasource=helm registryUrl=https://kubernetes.github.io/ingress-nginx depName=ingress-nginx versioning=semver
	))
}
