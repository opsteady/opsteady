package component

import (
	"github.com/opsteady/opsteady/cli/tasks"
)

// HelmChart contains information of the Helm chart
// values.yaml is expected to be in HELM_FOLDER/RELEASE/values.yaml
type HelmChart struct {
	Release   string
	Version   string
	Namespace string
}

// NewHelmChart creates a chart with namespace set to platform
// ValuesFileName wil be set later when using the HelmChart
func NewHelmChart(release, version string) *HelmChart {
	return &HelmChart{
		Namespace: "platform",
		Release:   release,
		Version:   version,
	}
}

// UseHelm initializes helm charts information
func (c *DefaultComponent) UseHelm(charts ...*HelmChart) {
	c.HelmCharts = charts
}

func (c *DefaultComponent) LoginToHelmRegistry() {
	c.Logger.Info().Msg("Preparing Helm environment...")

	// TODO: we should have a separate SP for accessing Helm registry
	mgmtCreds, err := c.Credentials.Azure("management")
	if err != nil {
		c.Logger.Fatal().Err(err).Msg("could not get management credentials to prepare helm")
	}

	c.Logger.Debug().Msg("Logging in to Helm repository")
	command := tasks.NewCommand("helm", c.GlobalConfig.TmpFolder)
	command.AddArgs(
		"registry",
		"login",
		c.GlobalConfig.ManagementHelmRepository,
		"--username",
		mgmtCreds["client_id"].(string),
		"--password",
		mgmtCreds["client_secret"].(string))
	command.AddEnv("HELM_EXPERIMENTAL_OCI", "1")

	if err := command.Run(); err != nil {
		c.Logger.Fatal().Err(err).Msg("could not login to helm registry")
	}
}
