// Split the list of the components into a separate package so there is no cyclic dependency
package components

import (
	capabilitiesCertificatesAWS "github.com/opsteady/opsteady/capabilities/certificates/aws/cicd"
	capabilitiesCertificatesAzure "github.com/opsteady/opsteady/capabilities/certificates/azure/cicd"
	capabilitiesCertificatesLocal "github.com/opsteady/opsteady/capabilities/certificates/local/cicd"
	capabilitiesDNSAWS "github.com/opsteady/opsteady/capabilities/dns/aws/cicd"
	capabilitiesDNSAzure "github.com/opsteady/opsteady/capabilities/dns/azure/cicd"
	capabilitiesDNSLocal "github.com/opsteady/opsteady/capabilities/dns/local/cicd"
	capabilitiesLoadbalancing "github.com/opsteady/opsteady/capabilities/loadbalancing/cicd"
	capabilitiesUserAuth "github.com/opsteady/opsteady/capabilities/user-auth/cicd"
	cli "github.com/opsteady/opsteady/cicd"
	"github.com/opsteady/opsteady/cli/component"
	dockerBase "github.com/opsteady/opsteady/docker/base/cicd"
	dockerCicd "github.com/opsteady/opsteady/docker/cicd/cicd"
	foundationAWS "github.com/opsteady/opsteady/foundation/aws/cicd"
	foundationAzure "github.com/opsteady/opsteady/foundation/azure/cicd"
	foundationLocal "github.com/opsteady/opsteady/foundation/local/cicd"
	kubernetesAWSCluster "github.com/opsteady/opsteady/kubernetes/aws/cluster/cicd"
	kubernetesAWSLoadbalancing "github.com/opsteady/opsteady/kubernetes/aws/loadbalancing/cicd"
	kubernetesAWSNetworkPolicies "github.com/opsteady/opsteady/kubernetes/aws/network-policies/cicd"
	kubernetesAWSStorageEBS "github.com/opsteady/opsteady/kubernetes/aws/storage/ebs/cicd"
	kubernetesAWSStorageEFS "github.com/opsteady/opsteady/kubernetes/aws/storage/efs/cicd"
	kubernetesAzureCluster "github.com/opsteady/opsteady/kubernetes/azure/cluster/cicd"
	kubernetesAzurePodIdentity "github.com/opsteady/opsteady/kubernetes/azure/pod-identity/cicd"
	kubernetesBootstrap "github.com/opsteady/opsteady/kubernetes/bootstrap/cicd"
	kubernetesLocalCluster "github.com/opsteady/opsteady/kubernetes/local/cluster/cicd"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	managementInfra "github.com/opsteady/opsteady/management/infra/cicd"
	managementVaultConfig "github.com/opsteady/opsteady/management/vault/config/cicd"
	managementVaultInfra "github.com/opsteady/opsteady/management/vault/infra/cicd"
)

func init() {
	Targets.addComponent(cli.Instance)
	Targets.addComponent(dockerBase.Instance)
	Targets.addComponent(dockerCicd.Instance)
	Targets.addComponent(managementBootstrap.Instance)
	Targets.addComponent(managementInfra.Instance)
	Targets.addComponent(managementVaultInfra.Instance)
	Targets.addComponent(managementVaultConfig.Instance)
	Targets.addComponent(foundationAzure.Instance)
	Targets.addComponent(foundationAWS.Instance)
	Targets.addComponent(kubernetesAWSCluster.Instance)
	Targets.addComponent(kubernetesAzureCluster.Instance)
	Targets.addComponent(kubernetesBootstrap.Instance)
	Targets.addComponent(kubernetesAzurePodIdentity.Instance)
	Targets.addComponent(foundationLocal.Instance)
	Targets.addComponent(kubernetesAWSStorageEBS.Instance)
	Targets.addComponent(kubernetesAWSStorageEFS.Instance)
	Targets.addComponent(kubernetesAWSNetworkPolicies.Instance)
	Targets.addComponent(kubernetesAWSLoadbalancing.Instance)
	Targets.addComponent(kubernetesLocalCluster.Instance)
	Targets.addComponent(capabilitiesDNSAzure.Instance)
	Targets.addComponent(capabilitiesDNSAWS.Instance)
	Targets.addComponent(capabilitiesDNSLocal.Instance)
	Targets.addComponent(capabilitiesCertificatesAWS.Instance)
	Targets.addComponent(capabilitiesCertificatesAzure.Instance)
	Targets.addComponent(capabilitiesCertificatesLocal.Instance)
	Targets.addComponent(capabilitiesLoadbalancing.Instance)
	Targets.addComponent(capabilitiesUserAuth.Instance)
}

// AllTargets contains all targets
type AllTargets struct {
	All []*Target
}

// Target contains sorted groups based on dependency for specific target
type Target struct {
	Name   component.Target
	Groups []*Group
}

// Group contains sorted components based on dependency for specific group
type Group struct {
	Name       component.Group
	Components []component.Component
}

// Targets contains a list of component initializers
var Targets = AllTargets{}

func (t *AllTargets) addComponent(c component.Component) {
	meta := c.GetMetadata()
	for _, ta := range meta.Targets {
		target := getTargetOrCreateEmptyTarget(t, ta)
		group := getGroupOrCreateEmptyGroup(target, meta.Group)
		group.Components = append(group.Components, c)
	}
}

func getTargetOrCreateEmptyTarget(t *AllTargets, target component.Target) *Target {
	for _, ta := range t.All {
		if target == ta.Name {
			return ta
		}
	}

	ta := &Target{
		Name: target,
	}
	t.All = append(t.All, ta)

	return ta
}

func getGroupOrCreateEmptyGroup(target *Target, group component.Group) *Group {
	for _, g := range target.Groups {
		if group == g.Name {
			return g
		}
	}

	ga := &Group{
		Name: group,
	}
	target.Groups = append(target.Groups, ga)

	return ga
}

// FindComponent searches for the component or returns nil
// We are accepting Target and Group to be empty, this however doesn't always work, if you are for example selecting bootstrap
func (t *AllTargets) FindComponent(ta component.Target, g component.Group, name string) component.Component {
	for _, tar := range t.All {
		if ta == tar.Name || string(ta) == "" {
			for _, ga := range tar.Groups {
				if g == ga.Name || string(g) == "" {
					for _, c := range ga.Components {
						if name == c.GetMetadata().Name {
							return c
						}
					}
				}
			}
		}
	}

	return nil
}
