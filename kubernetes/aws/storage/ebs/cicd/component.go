package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// KubernetesAWSStorageEBS is a component for the AWS load balancer controller.
type KubernetesAWSStorageEBS struct {
	component.DefaultComponent
}

var Instance = &KubernetesAWSStorageEBS{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "storage_ebs"
	m.Group = component.KubernetesAddons
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures KubernetesAWSStorageEBS before running
func (k *KubernetesAWSStorageEBS) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	k.SetVaultInfoToComponentConfig()
	k.UseHelm(component.NewHelmChart(
		"aws-ebs-csi-driver",
		"2.6.4", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/aws-ebs-csi-driver depName=aws-ebs-csi-driver versioning=semver
	))
}
