package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
)

// FoundationLocal is a component for the foundation
type FoundationLocal struct {
	component.DefaultComponent
}

var Instance = &FoundationLocal{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "foundation"
	m.Group = component.Foundation
	m.AddTarget(component.TargetLocal)
	Instance.Metadata = &m
}

// Configure configures FoundationLocal before running
func (f *FoundationLocal) Configure(defaultComponent component.DefaultComponent) {
	f.DefaultComponent = defaultComponent
	f.Terraform = "" // Use root of the folder
	f.AddRequiresInformationFrom(managementInfra.Instance.GetMetadata())
	f.AddManagementCredentialsToComponentConfig()
	f.SetVaultInfoToComponentConfig()
}
