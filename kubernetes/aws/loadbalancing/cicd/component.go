package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// KubernetesAWSLoadbalancing is a component for the AWS load balancer controller.
type KubernetesAWSLoadbalancing struct {
	component.DefaultComponent
}

var Instance = &KubernetesAWSLoadbalancing{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "loadbalancing"
	m.Group = component.KubernetesAddons
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures KubernetesAWSLoadbalancing before running
func (k *KubernetesAWSLoadbalancing) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	k.SetVaultInfoToComponentConfig()
	k.UseHelm(component.NewHelmChart(
		"aws-load-balancer-controller",
		"1.4.6", // renovate: datasource=helm registryUrl=https://aws.github.io/eks-charts depName=aws-load-balancer-controller versioning=semver
	))
}
