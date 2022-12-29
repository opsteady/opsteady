package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
	managementVaultInfra "github.com/opsteady/opsteady/management/vault/infra/cicd"
)

// ManagementBootstrap is a component for the management bootstrap
type ManagementVaultConfig struct {
	component.DefaultComponent
}

var Instance = &ManagementVaultConfig{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "config"
	m.Group = component.Vault
	m.AddTarget(component.TargetManagement)
	m.AddDependency(managementVaultInfra.Instance.GetMetadata())
	m.AddGroupDependency(component.Management)
	Instance.Metadata = &m
}

// Configure configures ManagementVaultConfig before running
func (m *ManagementVaultConfig) Configure(defaultComponent component.DefaultComponent) {
	m.DefaultComponent = defaultComponent
	m.Terraform = "" // Use root of the folder
	m.SetVaultInfoToComponentConfig()
	m.AddAzureADCredentialsToComponentConfig()
	m.AddRequiresInformationFrom(managementInfra.Instance.GetMetadata())
}
