package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAWSCluster is a component for creating Kubernetes (EKS)
type KubernetesAWSCluster struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAWSCluster struct
func (k *KubernetesAWSCluster) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.Terraform = "" // Use root of the folder
	k.DefaultComponent.RequiresComponents("foundation-aws")
}
