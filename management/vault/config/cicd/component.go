package cicd

import "github.com/opsteady/opsteady/cli/component"

// ManagementBootstrap is a component for the management bootstrap
type ManagementVaultConfig struct {
	component.DefaultComponent
}

// Initialize creates a new managementBootstrap struct
func (m *ManagementVaultConfig) Initialize(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
}
