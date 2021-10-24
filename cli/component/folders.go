package component

import (
	"errors"
	"fmt"
	"os"
)

// DetermineOrderOfExecution checks which folders exist (for example terraform, helm, etc..)
// and returns this as a list, thereby determining what will be executed.
// The default order is terraform, helm and then kubernetes but this can be
// overridden with for example OverrideDeployOrder.
func (c *DefaultComponent) DetermineOrderOfExecution() []string {
	folders := []string{}
	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.ComponentFolder, c.Terraform)); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Terraform)
	}
	if _, err := os.Stat(fmt.Sprintf("%s/%s", c.ComponentFolder, c.Helm)); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Helm)
	}
	if _, err := os.Stat(c.KubectlFolder()); !errors.Is(err, os.ErrNotExist) {
		folders = append(folders, c.Kubectl)
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

// KubectlFolder returns the kubectl folder inside the component including the component folder
func (c *DefaultComponent) KubectlFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.Kubectl)
}

// KubectlTmpFolder creates (if needed) and returns the kubectl folder inside the tmp folder including the component folder
func (c *DefaultComponent) KubectlTmpFolder() string {
	tmpFolder := fmt.Sprintf("%s/%s/%s", c.GlobalConfig.TmpFolder, c.ComponentFolder, c.Kubectl)
	if err := os.MkdirAll(tmpFolder, os.ModePerm); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not create kubectl tmp folder")
	}
	return tmpFolder
}

// DockerFolder returns the docker folder inside the component including the component folder
func (c *DefaultComponent) DockerFolder() string {
	return fmt.Sprintf("%s/%s", c.ComponentFolder, c.Docker)
}
