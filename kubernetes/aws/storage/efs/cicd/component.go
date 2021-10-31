package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAWSStorageEFS is a component for the AWS load balancer controller.
type KubernetesAWSStorageEFS struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAWSLoadbalancing component
func (k *KubernetesAWSStorageEFS) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
	k.DefaultComponent.UseHelm(&component.HelmChart{
		Release:   "aws-efs-csi-driver",
		Version:   "2.2.0",
		Namespace: "platform",
	})
}
