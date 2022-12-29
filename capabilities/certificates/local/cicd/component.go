package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationLocal "github.com/opsteady/opsteady/foundation/local/cicd"
	kubernetesLocalCluster "github.com/opsteady/opsteady/kubernetes/local/cluster/cicd"
)

// CapabilitiesCertificatesLocal is a local implementation for the certificates controller
type CapabilitiesCertificatesLocal struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesCertificatesLocal{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "certificates"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetLocal)
	m.AddGroupDependency(component.Kubernetes)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesCertificatesLocal before running
func (c *CapabilitiesCertificatesLocal) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationLocal.Instance.GetMetadata(), kubernetesLocalCluster.Instance.GetMetadata())
	c.AddAzureADCredentialsToComponentConfig()
	c.SetVaultInfoToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.10.1", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
