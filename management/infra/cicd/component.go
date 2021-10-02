package cicd

import "github.com/opsteady/opsteady/cli/component"

// ManagementInfra is a component for the management infrastructure
type ManagementInfra struct {
	component.DefaultComponent
}

// Initialize creates a new managementBootstrap struct
func (m *ManagementInfra) Initialize(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
}
