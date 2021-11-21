package cicd

import "github.com/opsteady/opsteady/cli/component"

// KubernetesAWSNetworkPolicies is a component for the AWS load balancer controller.
type KubernetesAWSNetworkPolicies struct {
	component.DefaultComponent
}

// Initialize creates a new KubernetesAWSLoadbalancing component
func (k *KubernetesAWSNetworkPolicies) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("foundation-aws", "kubernetes-aws-cluster")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
	k.DefaultComponent.UseHelm(component.NewHelmChart(
		"tigera-operator",
		"v3.21.0", // renovate: datasource=helm registryUrl=https://docs.projectcalico.org/charts depName=tigera-operator versioning=semver
	))
}
