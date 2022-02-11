package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
)

// ManagementBootstrap is a component for the management bootstrap
type ManagementBootstrap struct {
	component.DefaultComponent
}

var Instance = &ManagementBootstrap{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "bootstrap"
	m.Group = component.Management
	m.AddTarget(component.TargetManagement)
	Instance.Metadata = &m
}

// Configure configures ManagementBootstrap before running
func (m *ManagementBootstrap) Configure(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.Terraform = "" // Use root of the folder
}
