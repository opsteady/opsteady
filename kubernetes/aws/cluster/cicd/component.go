package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
)

// KubernetesAWSCluster is a component for creating Kubernetes (EKS)
type KubernetesAWSCluster struct {
	component.DefaultComponent
}

var Instance = &KubernetesAWSCluster{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "cluster"
	m.Group = component.Kubernetes
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.Foundation)
	Instance.Metadata = &m
}

// Configure configures KubernetesAWSCluster before running
func (k *KubernetesAWSCluster) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.Terraform = "" // Use root of the folder
	k.SetVaultInfoToComponentConfig()
	k.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata())
}
