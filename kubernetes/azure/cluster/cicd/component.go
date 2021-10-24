package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAzure is a component for creating Kubernetes (AKS)
type KubernetesAzure struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAzure struct
func (k *KubernetesAzure) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.Terraform = "" // Use root of the folder
	k.DefaultComponent.RequiresComponents("foundation-azure")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
}
