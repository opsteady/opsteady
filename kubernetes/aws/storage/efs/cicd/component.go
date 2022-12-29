package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// KubernetesAWSStorageEFS is a component for the AWS load balancer controller.
type KubernetesAWSStorageEFS struct {
	component.DefaultComponent
}

var Instance = &KubernetesAWSStorageEFS{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "storage_efs"
	m.Group = component.KubernetesAddons
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures KubernetesAWSStorageEFS before running
func (k *KubernetesAWSStorageEFS) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	k.SetVaultInfoToComponentConfig()
	k.UseHelm(component.NewHelmChart(
		"aws-efs-csi-driver",
		"2.3.5", // renovate: datasource=helm registryUrl=https://kubernetes-sigs.github.io/aws-efs-csi-driver depName=aws-efs-csi-driver versioning=semver
	))
}
