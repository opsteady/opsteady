package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
)

// KubernetesBootstrap is a component for bootstraping Kubernetes clusters
type KubernetesBootstrap struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesBootstrap struct
func (k *KubernetesBootstrap) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("management-infra", "foundation-aws", "foundation-azure", "foundation-local", "kubernetes-azure-cluster", "kubernetes-aws-cluster")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
	k.DefaultComponent.AddAzureADCredentialsToComponentConfig()
}
