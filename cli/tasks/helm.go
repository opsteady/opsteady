package tasks

import (
	"fmt"

	"github.com/rs/zerolog"
)

// Helm contains information to deploy the chart
type Helm struct {
	tmpFolder string
	logger    *zerolog.Logger
}

// NewHelm returns a Helm task runner
func NewHelm(tmpFolder string, logger *zerolog.Logger) *Helm {
	return &Helm{
		tmpFolder: tmpFolder,
		logger:    logger,
	}
}

// Upgrade installs or upgrades Helm releases
// TODO function not tested yet with OCI https://github.com/helm/helm/pull/9782
func (h *Helm) Upgrade(url, name, namespace, version string, dryRun bool) error {
	h.logger.Info().Str("release", url).Msg("Running Helm upgrade for release")

	command := NewCommand("helm", h.tmpFolder)
	command.AddArgs(
		"upgrade",
		"--install",
		"--atomic",
		fmt.Sprintf("--dry-run=%t", dryRun),
		"--namespace",
		namespace,
		"--version",
		version,
		url,
	)
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")

	return command.Run()
}

// Delete deletes the release
func (h *Helm) Delete(name, namespace string, dryRun bool) error {
	h.logger.Info().Str("release", name).Msg("Remove release")

	command := NewCommand("helm", h.tmpFolder)
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
