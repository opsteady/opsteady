package component

import (
	"errors"
	"fmt"
	"os"
)

// DetermineOrderOfExecution checks which folders exist (for example terraform, helm, etc..)
// and returns this as a list, thereby determining what will be executed and when.
// The default order is:
// terraform -> create cloud resources like IAM for pods and controllers to use
// crd -> we don't want to install CRDs from helm, having a separate step makes sense especially on destroy
// setup -> setup anything before installing helm, like Azure Pod identity
// helm -> install one or more helm charts
// post_setup -> same as setup, use for resources like Prometheus instances, important for destroy order
// docker -> create docker image, used in build step only
// this can be overridden with for example OverrideDeployOrder.
func (c *DefaultComponent) DetermineOrderOfExecution() []string {
	folders := []string{}
	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.ComponentFolder, c.Terraform)); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Terraform)
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.ComponentFolder, c.CRD)); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.CRD)
	}

	if _, err := os.Stat(c.KubeSetupFolder()); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.KubeSetup)
	}

	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.ComponentFolder, c.Helm)); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Helm)
	}

	if _, err := os.Stat(c.KubePostSetupFolder()); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.KubePostSetup)
	}

	if _, err := os.Stat(c.DockerFolder()); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Docker)
	}

	return folders
}

// TerraformFolder returns the terraform folder inside the component including the component folder
func (c *DefaultComponent) TerraformFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.Terraform)
}

// CRDFolder returns the CRD folder inside the component including the component folder
func (c *DefaultComponent) CRDFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.CRD)
}

// HelmFolder returns the helm folder inside the component including the component folder
func (c *DefaultComponent) HelmFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.Helm)
}

// HelmTmpFolder creates (if needed) and returns the helm folder inside the tmp folder including the component folder
func (c *DefaultComponent) HelmTmpFolder(chartName string) string {
	tmpFolder := fmt.Sprintf("%s/%s/%s/%s", c.GlobalConfig.TmpFolder, c.ComponentFolder, c.Helm, chartName)

	if err := os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not create helm tmp folder")
	}

	return tmpFolder
}

// KubeSetupFolder returns the KubeSetup folder inside the component including the component folder
func (c *DefaultComponent) KubeSetupFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.KubeSetup)
}

// KubeSetupTmpFolder creates (if needed) and returns the KubeSetup folder inside the tmp folder including the component folder
func (c *DefaultComponent) KubeSetupTmpFolder() string {
	tmpFolder := fmt.Sprintf("%s/%s/%s", c.GlobalConfig.TmpFolder, c.ComponentFolder, c.KubeSetup)

	if err := os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not create KubeSetup tmp folder")
	}

	return tmpFolder
}

// KubePostSetupFolder returns the KubePostSetup folder inside the component including the component folder
func (c *DefaultComponent) KubePostSetupFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.KubePostSetup)
}

// KubePostSetupTmpFolder creates (if needed) and returns the KubePostSetup folder inside the tmp folder including the component folder
func (c *DefaultComponent) KubePostSetupTmpFolder() string {
	tmpFolder := fmt.Sprintf("%s/%s/%s", c.GlobalConfig.TmpFolder, c.ComponentFolder, c.KubePostSetup)

	if err := os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not create PostSetup tmp folder")
	}

	return tmpFolder
}

// DockerFolder returns the docker folder inside the component including the component folder
func (c *DefaultComponent) DockerFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.Docker)
}
