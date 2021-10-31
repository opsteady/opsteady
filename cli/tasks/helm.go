package tasks

import (
	"fmt"

	"github.com/rs/zerolog"
)

// Helm contains information to deploy the chart
type Helm struct {
	logger *zerolog.Logger
}

// NewHelm returns a Helm task runner
func NewHelm(logger *zerolog.Logger) *Helm {
	return &Helm{
		logger: logger,
	}
}

// Upgrade installs or upgrades Helm releases
func (h *Helm) Upgrade(valuesFolder, url, name, namespace, version string, dryRun bool) error {
	h.logger.Info().Str("release", name).Msg("Running Helm upgrade for release")

	command := NewCommand("helm", valuesFolder)
	command.AddArgs(
		"upgrade",
		name,
		fmt.Sprintf("oci://%s/helm/%s", url, name),
		"--install",
		"--atomic",
		fmt.Sprintf("--dry-run=%t", dryRun),
		"--namespace",
		namespace,
		"--version",
		version,
		"--values",
		fmt.Sprintf("%s/values.yaml", valuesFolder),
	)
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")

	return command.Run()
}

// Delete deletes the release
func (h *Helm) Delete(valuesFolder, name, namespace string, dryRun bool) error {
	h.logger.Info().Str("release", name).Msg("Remove release")

	command := NewCommand("helm", valuesFolder)
	command.AddArgs(
		"uninstall",
		fmt.Sprintf("--dry-run=%t", dryRun),
		"--namespace",
		namespace,
		name,
	)

	return command.Run()
}

// Save accepts the root of the chart that needs to be saved, and the full url
func (h *Helm) Save(path, url string) error {
	h.logger.Info().Msg("Package the helm chart")

	command := NewCommand("helm", path)
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")
	command.AddArgs(
		"chart",
		"save",
		".",
		url,
	)

	return command.Run()
}

// Push pushes the chart to repository
func (h *Helm) Push(path, url string) error {
	h.logger.Info().Msg("Push the helm chart")

	command := NewCommand("helm", path)
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")
	command.AddArgs(
		"chart",
		"push",
		url,
	)

	return command.Run()
}

// LoginToHelmRegistry logs to a registry
func (h *Helm) LoginToHelmRegistry(user, pass, registry, tmpFolder string) error {
	h.logger.Debug().Msg("Logging in to Helm repository")
	command := NewCommand("helm", tmpFolder)
	command.AddArgs(
		"registry",
		"login",
		registry,
		"--username",
		user,
		"--password",
		pass)
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")

	return command.Run()
}
