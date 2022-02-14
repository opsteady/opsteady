package cicd

import "github.com/opsteady/opsteady/cli/component"

// Portal is a component for the Opsteady interface
type Portal struct {
	component.DefaultComponent
}

// Initialize creates a new Portal struct
func (p *Portal) Initialize(defaultComponent component.DefaultComponent) {
	p.DefaultComponent = defaultComponent
	p.Terraform = "" // Use root of the folder
	p.RequiresComponents("management-infra")
	p.AddManagementCredentialsToComponentConfig()
	p.SetVaultInfoToComponentConfig()
}
