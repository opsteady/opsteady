package tasks

import "github.com/rs/zerolog"

// Yaml contains information to run Yamllint command
type Yaml struct {
	logger *zerolog.Logger
}

// NewYaml returns a Yaml CLI task runner
func NewYaml(logger *zerolog.Logger) *Yaml {
	return &Yaml{
		logger: logger,
	}
}

// Lint lints the yaml
func (a *Yaml) Lint(folder string) error {
	a.logger.Info().Msg("Run yaml lint")

	command := NewCommand("yamllint", folder)
	command.AddArgs(".")

	return command.Run()
}
