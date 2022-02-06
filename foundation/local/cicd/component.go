package cicd

import "github.com/opsteady/opsteady/cli/component"

// FoundationLocal is a component for the foundation
type FoundationLocal struct {
	component.DefaultComponent
}

// Initialize creates a new FoundationLocal struct
func (f *FoundationLocal) Initialize(defaultComponent component.DefaultComponent) {
	f.DefaultComponent = defaultComponent
	f.DefaultComponent.Terraform = "" // Use root of the folder
	f.RequiresComponents("management-infra")
	f.DefaultComponent.AddManagementCredentialsToComponentConfig()
	f.DefaultComponent.SetVaultInfoToComponentConfig()
}
