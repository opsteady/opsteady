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
	k.DefaultComponent.UseHelm(component.NewHelmChart(
		"aws-load-balancer-controller",
		"1.4.0", // renovate: datasource=helm registryUrl=https://aws.github.io/eks-charts depName=aws-load-balancer-controller versioning=semver
	))
}
