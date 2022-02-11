package component

import "fmt"

// Metadata holds meta information of the component and it's dependencies when seting up the platform
type Metadata struct {
	Name                    string
	Description             string
	Group                   Group
	DependsOn               []*Metadata // Defines dependencies inside the group if needed
	DependsOnGroup          []Group     // Defines the group that needs to be finished first
	RequiresInformationFrom []*Metadata // Dependencies component has (used for fetching data from Vault)
	Targets                 []Target
}

// FullName returns the name inclusive targer, group and name
func (m *Metadata) FullName(e Target) string {
	return fmt.Sprintf("%s - %s - %s", e, m.Group, m.Name)
}

// VariableNames returns the name as used when reading variables from this component
func (m *Metadata) VariableNames(e Target) string {
	return fmt.Sprintf("%s_%s", e, m.Name)
}

// AddTarget adds targets to Metadata
func (m *Metadata) AddTarget(dep ...Target) {
	m.Targets = append(m.Targets, dep...)
}

// AddDependency ads dependencies to Metadata
func (m *Metadata) AddDependency(dep ...*Metadata) {
	m.DependsOn = append(m.DependsOn, dep...)
}

// AddGroupDependency ads group dependencies to Metadata
func (m *Metadata) AddGroupDependency(dep ...Group) {
	m.DependsOnGroup = append(m.DependsOnGroup, dep...)
}

// IsDependedOnGroup return true if the Metadata has a dependency on the group
func (m *Metadata) IsDependedOnGroup(c Group) bool {
	for _, v := range m.DependsOnGroup {
		if v == c {
			return true
		}
	}

	return false
}

// IsDependedOn return true if the Metadata has a dependency on the other Metadata component
func (m *Metadata) IsDependedOn(meta *Metadata) bool {
	for _, v := range m.DependsOn {
		if v == meta {
			return true
		}
	}

	return false
}

// DependsOnNames returns unique list of names that this component depends on
func (m *Metadata) DependsOnNames() []string {
	// Put in a map first to remove duplicate names
	depMapNames := make(map[string]bool)
	for _, dep := range m.DependsOn {
		depMapNames[dep.Name] = true
	}

	list := []string{}
	for name := range depMapNames {
		list = append(list, name)
	}

	return list
}

// AddRequiresInformationFrom ads Components which the component needs data from
func (m *Metadata) AddRequiresInformationFrom(dep ...*Metadata) {
	m.RequiresInformationFrom = append(m.RequiresInformationFrom, dep...)
}

type Group string
type Target string

const (
	Management        Group = "management"
	Vault             Group = "vault"
	Docker            Group = "docker"
	Cli               Group = "cli"
	Foundation        Group = "foundation"
	Kubernetes        Group = "kubernetes"
	KubernetesAddons  Group = "kubernetes_addons"
	CapabilitiesBasic Group = "capabilities_basic"
	CapabilitiesAuth  Group = "capabilities_auth"

	TargetDev        Target = "dev"
	TargetDocker     Target = "docker"
	TargetManagement Target = "management"
	TargetAzure      Target = "azure"
	TargetAws        Target = "aws"
	TargetLocal      Target = "local"
)

// DefaultMetadata should be used for initializing the Metadata so whenever we add
// new things to the object not all the code breaks at once, we can gradually add the desired information
func DefaultMetadata() Metadata {
	return Metadata{
		Name:                    "Change me",
		Description:             "",
		Targets:                 []Target{},
		Group:                   "Change me",
		DependsOn:               []*Metadata{},
		DependsOnGroup:          []Group{},
		RequiresInformationFrom: []*Metadata{},
	}
}

// AllTargets returns all targets as list
func AllTargets() []Target {
	return []Target{TargetManagement, TargetAzure, TargetAws, TargetLocal, TargetDev, TargetDocker}
}
