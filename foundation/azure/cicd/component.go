package cicd

import "github.com/opsteady/opsteady/cli/component"

// FoundationAzure is a component for the foundation
type FoundationAzure struct {
	component.DefaultComponent
}

// Initialize creates a new FoundationAzure struct
func (f *FoundationAzure) Initialize(defaultComponent component.DefaultComponent) {
	f.DefaultComponent = defaultComponent
	f.DefaultComponent.Terraform = "" // Use root of the folder
	f.RequiresComponents("management-infra")
	f.DefaultComponent.AddManagementCredentialsToComponentConfig()
	f.DefaultComponent.SetVaultInfoToComponentConfig()
}
