package component

import (
	"github.com/opsteady/opsteady/cli/tasks"
)

// Build generates an artifact for the component that can be released.
func (c *DefaultComponent) Build() {
	executeOrder := c.DetermineOrderOfExecution()
	if len(c.OverrideDeployOrder) != 0 {
		executeOrder = c.OverrideDeployOrder
	}

	for _, folder := range executeOrder {
		switch folder {
		case c.Docker:
			c.BuildDocker()
		}
	}
}

// BuildDocker builds the docker container
func (c *DefaultComponent) BuildDocker() {
	docker := tasks.NewDocker(c.Logger)

	if err := docker.Build(c.DockerFolder(), c.DockerBuildInfo.FullURL(c.GlobalConfig.ManagementDockerRegistry), c.DockerBuildInfo.BuildArgs); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not build Docker container")
	}
}
