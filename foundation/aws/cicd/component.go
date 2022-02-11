package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
)

// FoundationAzure is a component for the foundation
type FoundationAWS struct {
	component.DefaultComponent
}

var Instance = &FoundationAWS{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "foundation"
	m.Group = component.Foundation
	m.AddTarget(component.TargetAws)
	Instance.Metadata = &m
}

// Configure configures FoundationAWS before running
func (f *FoundationAWS) Configure(defaultComponent component.DefaultComponent) {
	f.DefaultComponent = defaultComponent
	f.Terraform = "" // Use root of the folder
	f.AddManagementCredentialsToComponentConfig()
	f.SetVaultInfoToComponentConfig()
	f.AddRequiresInformationFrom(managementInfra.Instance.GetMetadata())
}
