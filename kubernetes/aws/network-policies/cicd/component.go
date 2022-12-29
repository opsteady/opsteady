package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// KubernetesAWSNetworkPolicies is a component for the AWS load balancer controller.
type KubernetesAWSNetworkPolicies struct {
	component.DefaultComponent
}

var Instance = &KubernetesAWSNetworkPolicies{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "network_policies"
	m.Group = component.KubernetesAddons
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures KubernetesAWSNetworkPolicies before running
func (k *KubernetesAWSNetworkPolicies) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	k.SetVaultInfoToComponentConfig()
	k.UseHelm(component.NewHelmChart(
		"tigera-operator",
		"v3.22.2", // renovate: datasource=helm registryUrl=https://docs.projectcalico.org/charts depName=tigera-operator versioning=semver
	))
}
