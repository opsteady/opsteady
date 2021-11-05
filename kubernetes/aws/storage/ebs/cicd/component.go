package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAWSStorageEBS is a component for the AWS load balancer controller.
type KubernetesAWSStorageEBS struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAWSLoadbalancing component
func (k *KubernetesAWSStorageEBS) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
	k.DefaultComponent.UseHelm(&component.HelmChart{
		Release:   "aws-ebs-csi-driver",
		Version:   "2.4.0", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/aws-ebs-csi-driver depName=aws-ebs-csi-driver versioning=semver
		Namespace: "platform",
	})
}
