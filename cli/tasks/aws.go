package tasks

import "github.com/rs/zerolog"

// Aws contains information to run AWS commands
type Aws struct {
	tmpFolder string
	logger    *zerolog.Logger
}

// NewAws returns a AWS CLI task runner
func NewAws(tmpFolder string, logger *zerolog.Logger) *Aws {
	return &Aws{
		tmpFolder: tmpFolder,
		logger:    logger,
	}
}

// LoginToEKS logs in to EKS
func (a *Aws) LoginToEKS(region, cluster string) error {
	a.logger.Info().Msg("Preparing EKS environment...")

	command := NewCommand("aws", a.tmpFolder)
	command.AddArgs("eks",
		"--region", region,
		"update-kubeconfig",
		"--name", cluster)

	return command.Run()
}
