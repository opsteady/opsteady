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

// func GetOrderedComponents() []SortedGroupsTree {
// 	allTargets := component.AllTargets()

// 	all := []SortedGroupsTree{}

// 	// Loop trough all sorted envs
// 	for _, target := range allTargets {
// 		groups := SortedGroupsTree{
// 			Target: target,
// 			Groups: []SortedComponentsTree{},
// 		}

// 		var firstGroup component.Group

// 		// Find the group that has no dependency as that is the starting point
// 		for _, comps := range Components[target] {
// 			for _, c := range comps {
// 				if len(c.GetMetadata().DependsOnGroup) == 0 {
// 					firstGroup = c.GetMetadata().Group
// 				}
// 			}
// 		}

// 		// Find the group dependency tree
// 		for _, group := range findGroupsDependingOnGroup(Components[target], []component.Group{firstGroup}, firstGroup) {
// 			components := Components[target][group]
// 			componentNames := []string{}

// 			for c := range components {
// 				componentNames = append(componentNames, c)
// 			}

// 			// Sort by name so we have a consistent order
// 			sort.Strings(componentNames)

// 			// Order the components inside a group based on the dependency
// 			orderedComponents := []component.ComponentInterface{}
// 			for _, cnOne := range componentNames {
// 				cOne := Components[target][group][cnOne]
// 				if len(cOne.GetMetadata().DependsOn) == 0 {
// 					orderedComponents = append(orderedComponents, cOne)
// 					for _, cnTwo := range componentNames {
// 						cTwo := Components[target][group][cnTwo]
// 						if cTwo.GetMetadata().IsDependedOn(cOne.GetMetadata()) {
// 							orderedComponents = append(orderedComponents, cTwo)
// 						}
// 					}
// 				}
// 			}

// 			// Add the group to the group object
// 			groups.Groups = append(groups.Groups, SortedComponentsTree{
// 				Group:      group,
// 				Components: orderedComponents,
// 			})
// 		}
// 		// Add the groups to the env object
// 		all = append(all, groups)
// 	}

// 	return all
// }

// func findGroupsDependingOnGroup(groups map[component.Group]map[string]component.ComponentInterface, list []component.Group, dependsOnGroup component.Group) []component.Group {
// 	for _, compsInGroup := range groups {
// 		for _, c := range compsInGroup {
// 			if c.GetMetadata().IsDependedOnGroup(dependsOnGroup) {
// 				// We are looping trough all components, so only add the group if it isn't already in the list
// 				if !componentGroupAlreadyInThelist(list, c.GetMetadata().Group) {
// 					list = append(list, c.GetMetadata().Group)
// 				}
// 				return findGroupsDependingOnGroup(groups, list, c.GetMetadata().Group)
// 			}
// 		}
// 	}
// 	return list
// }

// func componentGroupAlreadyInThelist(list []component.Group, c component.Group) bool {
// 	for _, v := range list {
// 		if v == c {
// 			return true
// 		}
// 	}

// 	return false
// }
