package cicd

import "github.com/opsteady/opsteady/cli/component"

// ManagementBootstrap is a component for the management bootstrap
type ManagementVaultInfra struct {
	component.DefaultComponent
}

// Initialize creates a new managementBootstrap struct
func (m *ManagementVaultInfra) Initialize(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.DefaultComponent.Terraform = "" // Use root of the folder
	m.DefaultComponent.AddAzureADCredentialsToComponentConfig()
	m.DefaultComponent.RequiresComponents("management-infra")
}
