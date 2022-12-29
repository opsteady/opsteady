package cicd

import (
	"github.com/opsteady/opsteady/cli/component"
	"github.com/opsteady/opsteady/cli/tasks"
)

// OpsteadyCli is a component for creating Opsteady CLI
type OpsteadyCli struct {
	component.DefaultComponent
}

var Instance = &OpsteadyCli{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "cli"
	m.Group = component.Cli
	m.AddTarget(component.TargetDev)
	Instance.Metadata = &m
}

// Configure configures OpsteadyCli before running
func (o *OpsteadyCli) Configure(defaultComponent component.DefaultComponent) {
	o.DefaultComponent = defaultComponent
}

func (o *OpsteadyCli) Validate() {
	lint := tasks.NewCommand("golangci-lint", o.ComponentFolder)
	lint.AddArgs("run", "--timeout", "10m")

	if err := lint.Run(); err != nil {
		o.Logger.Fatal().Err(err).Msg("Golang linting failed")
	}

	yaml := tasks.NewYaml(o.Logger)

	if err := yaml.Lint("."); err != nil {
		o.Logger.Fatal().Err(err).Msg("YAML linting failed")
	}
}

func (o *OpsteadyCli) Deploy() {
	o.Logger.Info().Msg("Deploy not supported for this component")
}

func (o *OpsteadyCli) Destroy() {
	o.Logger.Info().Msg("Destroy not supported for this component")
}

func (o *OpsteadyCli) Build() {
	o.Logger.Info().Msg("Build not supported for this component")
}

func (o *OpsteadyCli) Test() {
	o.Logger.Info().Msg("Test not supported for this component")
}

func (o *OpsteadyCli) Publish() {
	o.Logger.Info().Msg("Publish not supported for this component")
}
