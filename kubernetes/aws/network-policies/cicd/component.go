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
	k.DefaultComponent.UseHelm(&component.HelmChart{
		Release:   "tigera-operator",
		Version:   "v3.20.2",  // renovate: datasource=helm registryUrl=https://docs.projectcalico.org/charts depName=tigera-operator versioning=semver
		Namespace: "platform", // Everything is installed in the tigera-operator/calico-system namespaces. This is hardcoded in the chart.
	})
}
