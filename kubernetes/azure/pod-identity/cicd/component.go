package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// KubernetesAzurePodIdentity is a component for creating pod identity in Azure
type KubernetesAzurePodIdentity struct {
	component.DefaultComponent
}

var Instance = &KubernetesAzurePodIdentity{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "pod_identity"
	m.Group = component.KubernetesAddons
	m.AddTarget(component.TargetAzure)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures KubernetesAzurePodIdentity before running
func (k *KubernetesAzurePodIdentity) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.AddRequiresInformationFrom(foundationAzure.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata())
	k.UseHelm(component.NewHelmChart(
		"aad-pod-identity",
		"4.1.8", // renovate: datasource=helm registryUrl=https://raw.githubusercontent.com/Azure/aad-pod-identity/master/charts depName=aad-pod-identity versioning=semver
	))
}
