package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	"github.com/opsteady/opsteady/cli/tasks"
)

// OpsteadyCli is a component for creating Opsteady CLI
type OpsteadyCli struct {
	component.DefaultComponent
}

// Initialize creates a new OpsteadyCli struct
func (o *OpsteadyCli) Initialize(defaultComponent component.DefaultComponent) {
	o.DefaultComponent = defaultComponent
}

func (o *OpsteadyCli) Validate() {
	lint := tasks.NewCommand("golangci-lint", o.ComponentFolder)
	lint.AddArgs("run", "--timeout", "10m")
	if err := lint.Run(); err != nil {
		o.Logger.Fatal().Err(err).Msg("Linting failed")
	}
}

func (c *OpsteadyCli) Deploy() {
	c.Logger.Info().Msg("Deploy not supported for this component")
}

func (c *OpsteadyCli) Destroy() {
	c.Logger.Info().Msg("Destroy not supported for this component")
}

func (c *OpsteadyCli) Build() {
	c.Logger.Info().Msg("Build not supported for this component")
}

func (c *OpsteadyCli) Test() {
	c.Logger.Info().Msg("Test not supported for this component")
}

func (c *OpsteadyCli) Publish() {
	c.Logger.Info().Msg("Publish not supported for this component")
}
