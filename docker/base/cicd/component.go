package cicd

import "github.com/opsteady/opsteady/cli/component"

// DockerBase is a component for creating base Docker image
type DockerBase struct {
	component.DefaultComponent
}

var Instance = &DockerBase{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "base"
	m.Group = component.Docker
	m.AddTarget(component.TargetDocker)
	Instance.Metadata = &m
}

// Configure configures DockerBase before running
func (d *DockerBase) Configure(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	d.SetDockerBuildInfo("base", "2.0.0", nil)
}

func (d *DockerBase) Deploy() {
	d.Logger.Info().Msg("Deploy not supported for this component")
}

func (d *DockerBase) Destroy() {
	d.Logger.Info().Msg("Destroy not supported for this component")
}
