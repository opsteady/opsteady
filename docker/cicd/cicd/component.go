package cicd

import (
	"fmt"

	"github.com/opsteady/opsteady/cli/component"
)

// DockerCicd is a component for creating base Docker image
type DockerCicd struct {
	component.DefaultComponent
}

// Initialize creates a new DockerCicd struct
func (d *DockerCicd) Initialize(defaultComponent component.DefaultComponent) {
	d.DefaultComponent = defaultComponent
	d.Docker = "" // Use root of the folder
	buildArgs := map[string]string{
		"FROM_IMAGE": fmt.Sprintf("%s/%s:%s",
			d.GlobalConfig.ManagementDockerRegistry,
			"base",
			"2.0.0", // renovate: datasource=docker registryUrl=opsteadyos.azurecr.io depName=opsteadyos.azurecr.io/base versioning=semver
		),
		"VAULT_CA_STORAGE_ACCOUNT": d.GlobalConfig.VaultCaStorageAccountName,
	}
	d.SetDockerBuildInfo("cicd", "2.9.0", buildArgs)
}

func (d *DockerCicd) Deploy() {
	d.Logger.Info().Msg("Deploy not supported for this component")
}

func (d *DockerCicd) Destroy() {
	d.Logger.Info().Msg("Destroy not supported for this component")
}
