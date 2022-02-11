package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
)

// KubernetesAzure is a component for creating Kubernetes (AKS)
type KubernetesAzure struct {
	component.DefaultComponent
}

var Instance = &KubernetesAzure{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "cluster"
	m.Group = component.Kubernetes
	m.AddTarget(component.TargetAzure)
	m.AddGroupDependency(component.Foundation)
	Instance.Metadata = &m
}

// Configure configures KubernetesAzure before running
func (k *KubernetesAzure) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.Terraform = "" // Use root of the folder
	k.AddRequiresInformationFrom(foundationAzure.Instance.Metadata)
	k.SetVaultInfoToComponentConfig()
}
