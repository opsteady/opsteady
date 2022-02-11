package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAzurePodIdentity is a component for creating pod identity in Azure
type KubernetesAzurePodIdentity struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAzurePodIdentity struct
func (k *KubernetesAzurePodIdentity) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.RequiresComponents("foundation-azure", "kubernetes-azure-cluster")
	k.DefaultComponent.UseHelm(component.NewHelmChart(
		"aad-pod-identity",
		"4.1.8", // renovate: datasource=helm registryUrl=https://raw.githubusercontent.com/Azure/aad-pod-identity/master/charts depName=aad-pod-identity versioning=semver
	))
}

func (k *KubernetesAzurePodIdentity) Info() component.ComponentDepInfo {
	return component.ComponentDepInfo{
		Name:           "kubernetes-azure-pod-identity",
		Description:    "Ads AAD pod identity to AKS cluster",
		Group:          "Kubernetes Addons",
		DependsOn:      []string{"kubernetes-azure-cluster"},
		DependsOnGroup: "Kubernetes",
	}
}
