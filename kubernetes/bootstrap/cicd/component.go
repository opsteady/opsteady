package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	foundationLocal "github.com/opsteady/opsteady/foundation/local/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
)

// KubernetesBootstrap is a component for bootstraping Kubernetes clusters
type KubernetesBootstrap struct {
	component.DefaultComponent
}

var Instance = &KubernetesBootstrap{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "bootstrap"
	m.Group = component.Kubernetes
	m.AddTarget(component.TargetAzure, component.TargetAws)
	m.AddGroupDependency(component.Kubernetes)
	m.AddDependency(kubernetesAWSCluster.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata())
	Instance.Metadata = &m
}

// Configure configures KubernetesBootstrap before running
func (k *KubernetesBootstrap) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.SetVaultInfoToComponentConfig()
	k.AddAzureADCredentialsToComponentConfig()
	k.AddRequiresInformationFrom(
		managementInfra.Instance.GetMetadata(),
		foundationAWS.Instance.GetMetadata(),
		foundationAzure.Instance.GetMetadata(),
		foundationLocal.Instance.GetMetadata(),
		kubernetesAWSCluster.Instance.GetMetadata(),
		kubernetesAzureCluster.Instance.GetMetadata())
}
