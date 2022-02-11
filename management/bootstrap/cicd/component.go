package cicd

import "github.com/opsteady/opsteady/cli/component"

// ManagementBootstrap is a component for the management bootstrap
type ManagementBootstrap struct {
	component.DefaultComponent
}

// Initialize creates a new managementBootstrap struct
func (m *ManagementBootstrap) Initialize(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.DefaultComponent.Terraform = "" // Use root of the folder
}

var ComponentInfo = component.DefaultComponentInfoImpl()

func init() {
	ComponentInfo.Name = "management-bootstrap"
	ComponentInfo.Group = component.ManagementBootstrap
}

// Initialize creates a new managementBootstrap struct
func (c component.ComponentInfoImpl) Initialize(defaultComponent component.DefaultComponent) component.Component {
	m := &ManagementBootstrap{}
	m.DefaultComponent = defaultComponent
	m.DefaultComponent.Terraform = "" // Use root of the folder
	return m
}
