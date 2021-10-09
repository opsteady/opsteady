package cicd

import "github.com/opsteady/opsteady/cli/component"

// FoundationAzure is a component for the foundation
type FoundationAWS struct {
	component.DefaultComponent
}

// Initialize creates a new FoundationAzure struct
func (f *FoundationAWS) Initialize(defaultComponent component.DefaultComponent) {
	f.DefaultComponent = defaultComponent
	f.DefaultComponent.Terraform = "" // Use root of the folder
	f.RequiresComponents("management-infra")
	f.DefaultComponent.AddManagementCredentialsToComponentConfig()
	f.DefaultComponent.SetVaultInfoToComponentConfig()
}
