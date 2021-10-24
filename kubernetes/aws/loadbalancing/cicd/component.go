package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAWSLoadbalancing is a component for the AWS load balancer controller.
type KubernetesAWSLoadbalancing struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAWSLoadbalancing component
func (k *KubernetesAWSLoadbalancing) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
	k.DefaultComponent.UseHelm(&component.HelmChart{
		Release:   "aws-load-balancer-controller",
		Version:   "1.3.1",
		Namespace: "platform",
	})
}
