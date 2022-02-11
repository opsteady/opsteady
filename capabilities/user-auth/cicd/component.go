package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// UserAuth is an implementation for the user authentication
type UserAuth struct {
	component.DefaultComponent
}

var Instance = &UserAuth{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "user_auth"
	m.Group = component.CapabilitiesAuth
	m.AddTarget(component.TargetLocal, component.TargetAws, component.TargetAzure)
	m.AddGroupDependency(component.CapabilitiesBasic)
	Instance.Metadata = &m
}

// Configure configures UserAuth before running
func (c *UserAuth) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(
		foundationAWS.Instance.GetMetadata(),
		foundationAzure.Instance.GetMetadata(),
		kubernetesAzureCluster.Instance.GetMetadata(),
		kubernetesAWSCluster.Instance.GetMetadata())
	c.SetVaultInfoToComponentConfig()
	c.AddManagementCredentialsToComponentConfig()
	c.AddAzureADCredentialsToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"dex",
		"0.6.5", // renovate: datasource=helm registryUrl=https://charts.dexidp.io depName=dex versioning=semver
	))
}
