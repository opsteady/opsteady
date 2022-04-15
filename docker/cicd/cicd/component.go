package cicd

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/component"
)

// DockerCicd is a component for creating base Docker image
type DockerCicd struct {
	component.DefaultComponent
}

var Instance = &DockerCicd{}

func init() {
	m := component.DefaultMetadata()
	m.Name = "cicd"
	m.Group = component.Docker
	m.AddTarget(component.TargetDocker)
	Instance.Metadata = &m
}

// Configure configures DockerCicd before running
func (d *DockerCicd) Configure(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	buildArgs := map[string]string{
		"FROM_IMAGE": fmt.Sprintf("%s/%s:%s",
			d.GlobalConfig.ManagementDockerRegistry,
			"base",
			"1.0.0", // renovate: datasource=docker registryUrl=opsteadyos.azurecr.io depName=opsteadyos.azurecr.io/base versioning=semver
		),
		"VAULT_CA_STORAGE_ACCOUNT": d.GlobalConfig.VaultCaStorageAccountName,
	}
	d.SetDockerBuildInfo("cicd", "1.6.1", buildArgs)
}

func (d *DockerCicd) Deploy() {
	d.Logger.Info().Msg("Deploy not supported for this component")
}

func (d *DockerCicd) Destroy() {
	d.Logger.Info().Msg("Destroy not supported for this component")
}
