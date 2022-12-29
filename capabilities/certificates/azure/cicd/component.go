package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
)

// CapabilitiesCertificatesAzure is an Azure implementation for the certificates controller
type CapabilitiesCertificatesAzure struct {
	component.DefaultComponent
}

var Instance = &CapabilitiesCertificatesAzure{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "certificates"
	m.Group = component.CapabilitiesBasic
	m.AddTarget(component.TargetAzure)
	m.AddGroupDependency(component.KubernetesAddons)
	Instance.Metadata = &m
}

// Configure configures CapabilitiesCertificatesAzure before running
func (c *CapabilitiesCertificatesAzure) Configure(defaultComponent component.DefaultComponent) {
	c.DefaultComponent = defaultComponent
	c.AddRequiresInformationFrom(foundationAzure.Instance.GetMetadata(), kubernetesAzureCluster.Instance.GetMetadata())
	c.SetVaultInfoToComponentConfig()
	c.UseHelm(component.NewHelmChart(
		"cert-manager",
		"v1.10.1", // renovate: datasource=helm registryUrl=https://charts.jetstack.io depName=cert-manager versioning=semver
	))
}
