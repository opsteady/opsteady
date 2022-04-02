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
	k.DefaultComponent.UseHelm(component.NewHelmChart(
		"aws-efs-csi-driver",
		"2.2.4", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/aws-efs-csi-driver depName=aws-efs-csi-driver versioning=semver
	))
}
