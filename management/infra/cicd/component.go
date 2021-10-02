package cicd

import "github.com/opsteady/opsteady/cli/component"

// ManagementInfra is a component for the management infrastructure
type ManagementInfra struct {
	component.DefaultComponent
}

// Initialize creates a new managementInfra struct
func (m *ManagementInfra) Initialize(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.DefaultComponent.Terraform = "" // Use root of the folder
	m.DefaultComponent.AddAzureADCredentialsToComponentConfig()
}
