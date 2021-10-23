package tasks

import (
	"github.com/rs/zerolog"
)

// Kubectl contains information to deploy the Kubernetes manifest files
type Kubectl struct {
	logger *zerolog.Logger
}

// NewKubectl returns a Kubectl task runner
func NewKubectl(logger *zerolog.Logger) *Kubectl {
	return &Kubectl{
		logger: logger,
	}
}

// Apply applies the Kubernetes manifest files
func (k *Kubectl) Apply(manifestFolder string, dryRun bool) error {
	k.logger.Info().Str("manifestFolder", manifestFolder).Msg("Apply manifest files")

	command := NewCommand("kubectl", manifestFolder)
	command.AddArgs("apply")
	if dryRun {
		command.AddArgs("--dry-run=client", "-oyaml")
	}
	command.AddArgs("-f", ".")

	return command.Run()
}

// Delete deletes the Kubernetes manifest files
func (k *Kubectl) Delete(manifestFolder string, dryRun bool) error {
	k.logger.Info().Str("manifestFolder", manifestFolder).Msg("Remove manifest files")

	command := NewCommand("kubectl", manifestFolder)
	command.AddArgs("delete")
	if dryRun {
		command.AddArgs("--dry-run=client", "-oyaml")
	}
	command.AddArgs("-f", ".")

	return command.Run()
}
