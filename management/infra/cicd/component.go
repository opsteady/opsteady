package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
)

// ManagementInfra is a component for the management infrastructure
type ManagementInfra struct {
	component.DefaultComponent
}

var Instance = &ManagementInfra{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "infra"
	m.Group = component.Management
	m.AddTarget(component.TargetManagement)
	m.AddDependency(managementBootstrap.Instance.GetMetadata())
	Instance.Metadata = &m
}

// Configure configures ManagementInfra before running
func (m *ManagementInfra) Configure(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.Terraform = "" // Use root of the folder
	m.AddAzureADCredentialsToComponentConfig()
}
