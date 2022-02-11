package cicd

import (
	"os"

	"github.com/opsteady/opsteady/cli/component"
	"github.com/opsteady/opsteady/cli/tasks"
)

const (
	readWriteExecute = 0700
	storageFolder    = "/tmp/k3d/storage"
)

// KubernetesLocal is a component for creating Kubernetes (AKS)
type KubernetesLocal struct {
	component.DefaultComponent
	k3d *tasks.K3d
}

var Instance = &KubernetesLocal{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "cluster"
	m.Group = component.Kubernetes
	m.AddTarget(component.TargetLocal)
	m.AddGroupDependency(component.Foundation)
	Instance.Metadata = &m
}

// Configure configures KubernetesLocal before running
func (k *KubernetesLocal) Configure(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.SetVaultInfoToComponentConfig()
	k.k3d = tasks.NewK3d(k.ComponentFolder, k.Logger)
}

func (k *KubernetesLocal) Deploy() {
	k.SetPlatformInfoToComponentConfig()

	if err := os.MkdirAll(storageFolder, readWriteExecute); err != nil {
		k.Logger.Fatal().Err(err).Str("dir", storageFolder).Msg("could not create storage directory")
	}

	// Everything in the server folder will be applied on the server on boot
	if err := k.k3d.CreateCluster("k3d-config.yaml"); err != nil {
		k.Logger.Fatal().Err(err).Msg("could not create cluster locally")
	}
}

func (k *KubernetesLocal) Destroy() {
	k.SetPlatformInfoToComponentConfig()

	if err := k.k3d.DeleteCluster("opsteady"); err != nil {
		k.Logger.Fatal().Err(err).Msg("could not delete cluster locally")
	}
}
