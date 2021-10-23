package tasks

import "github.com/rs/zerolog"

type Docker struct {
	logger *zerolog.Logger
}

// NewDocker creates a Docker struct
func NewDocker(logger *zerolog.Logger) *Docker {
	return &Docker{
		logger: logger,
	}
}

func (d *Docker) Build(workingDir, fullImageName string) error {
	d.logger.Info().Msg("Running docker build")
	command := NewCommand("docker", workingDir)
	command.AddArgs("build", "-t", fullImageName, ".")
	return command.Run()
}

func (d *Docker) Push(workingDir, fullImageName string) error {
	d.logger.Info().Msg("Push docker image")
	command := NewCommand("docker", workingDir)
	command.AddArgs("push", fullImageName)
	return command.Run()
}
