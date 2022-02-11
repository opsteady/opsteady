package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
)

// CapabilitiesCertificatesAWS is an AWS implementation for the certificates controller
type CapabilitiesCertificatesAWS struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesCertificatesAWS{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "certificates"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetAws)
	m.AddGroupDependency(component.KubernetesAddons)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesCertificatesAWS before running
func (c *CapabilitiesCertificatesAWS) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.SetVaultInfoToComponentConfig()
	c.AddManagementCredentialsToComponentConfig()
	c.AddRequiresInformationFrom(foundationAWS.Instance.GetMetadata(), kubernetesAWSCluster.Instance.GetMetadata())
	c.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.7.1", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
