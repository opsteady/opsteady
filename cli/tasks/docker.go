package tasks

import (
	"fmt"

	"github.com/rs/zerolog"
)

type Docker struct {
	logger *zerolog.Logger
}

// NewDocker creates a Docker struct
func NewDocker(logger *zerolog.Logger) *Docker {
	return &Docker{
		logger: logger,
	}
}

// Build runs the `docker build` command.
func (d *Docker) Build(workingDir, fullImageName string, args map[string]string) error {
	d.logger.Info().Msg("Running docker build")

	command := NewCommand("docker", workingDir)
	command.AddArgs("build", "-t", fullImageName, ".")

	for k, v := range args {
		command.AddArgs("--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	return command.Run()
}

// Push runs the `docker push` command.
func (d *Docker) Push(workingDir, fullImageName string) error {
	d.logger.Info().Msg("Push docker image")

	command := NewCommand("docker", workingDir)
	command.AddArgs("push", fullImageName)

	return command.Run()
}

// Validate runs the `hadolint` command to validate the Dockerfile.
func (d *Docker) Validate(workingDir string) error {
	d.logger.Info().Msg("Validate docker image")

	command := NewCommand("hadolint", workingDir)
	command.AddArgs("--failure-threshold", "error", "Dockerfile")

	return command.Run()
}
