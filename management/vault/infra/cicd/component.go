package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
)

// ManagementBootstrap is a component for the management bootstrap
type ManagementVaultInfra struct {
	component.DefaultComponent
}

var Instance = &ManagementVaultInfra{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "infra"
	m.Group = component.Vault
	m.AddTarget(component.TargetManagement)
	m.AddGroupDependency(component.Management)
	Instance.Metadata = &m
}

// Configure configures ManagementVaultInfra before running
func (m *ManagementVaultInfra) Configure(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.Terraform = "" // Use root of the folder
	m.AddAzureADCredentialsToComponentConfig()
	m.AddRequiresInformationFrom(managementInfra.Instance.GetMetadata())
}
