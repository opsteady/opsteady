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
	k.DefaultComponent.UseHelm(component.NewHelmChart("aad-pod-identity", "4.1.6"))
}
