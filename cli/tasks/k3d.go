package tasks

import (
	"github.com/rs/zerolog"
)

// K3d contains information to run k3d commands
type K3d struct {
	tmpFolder string
	logger    *zerolog.Logger
}

// NewK3d returns a K3d task runner
func NewK3d(tmpFolder string, logger *zerolog.Logger) *K3d {

	return &K3d{
		tmpFolder: tmpFolder,
		logger:    logger,
	}
}

// CreateCluster creates the cluster
func (k K3d) CreateCluster(config string) error {
	k.logger.Info().Str("config", config).Msg("Create k3d cluster")
	command := NewCommand("k3d", k.tmpFolder)
	command.AddArgs("cluster", "create", "--config", config)

	return command.Run()
}

// DeleteCluster creates the cluster
func (k K3d) DeleteCluster(name string) error {
	k.logger.Info().Str("cluster", name).Msg("Create k3d cluster")

	command := NewCommand("k3d", k.tmpFolder)
	command.AddArgs("cluster", "delete", name)

	return command.Run()
}

// LoginToKubernetes creates the cluster
func (k K3d) LoginToKubernetes(name string) error {
	k.logger.Info().Str("cluster", name).Msg("Login to k3d cluster")

	command := NewCommand("k3d", k.tmpFolder)
	command.AddArgs("kubeconfig", "merge", "--kubeconfig-merge-default", "--kubeconfig-switch-context", name)

	return command.Run()
}
