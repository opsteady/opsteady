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

// Initialize creates a new KubernetesLocal struct
func (k *KubernetesLocal) Initialize(defaultComponent component.DefaultComponent) {
	k.DefaultComponent = defaultComponent
	k.DefaultComponent.RequiresComponents("foundation-azure")
	k.DefaultComponent.SetVaultInfoToComponentConfig()
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
