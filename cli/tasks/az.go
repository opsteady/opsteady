package tasks

import (
	"github.com/rs/zerolog"
)

// Az contains information to run Azure commands
type Az struct {
	tmpFolder string
	logger    *zerolog.Logger
}

// NewAz returns a az CLI task runner
func NewAz(tmpFolder string, logger *zerolog.Logger) *Az {
	return &Az{
		tmpFolder: tmpFolder,
		logger:    logger,
	}
}

// LoginToAzure logs in to Azure
func (a *Az) LoginToAzure(clientID, clientSecret, tenantID string) error {
	a.logger.Info().Msg("Log in to Azure environment...")

	command := NewCommand("az", a.tmpFolder)
	command.AddArgs(
		"login",
		"--service-principal",
		"-u", clientID,
		"-p", clientSecret,
		"--tenant", tenantID)

	return command.Run()
}

// LoginToAKS logs in to AKS
func (a *Az) LoginToAKS(clusterName, clusterResourceGroup string) error {
	command := NewCommand("az", a.tmpFolder)
	command.AddArgs(
		"aks",
		"get-credentials",
		"-n", clusterName,
		"-g", clusterResourceGroup,
		"--overwrite-existing")

	return command.Run()
}

// LoginToAcr logs in to ACR
func (a *Az) LoginToAcr(acrName, subscription, clientID, clientSecret string) error {
	a.logger.Info().Str("acr", acrName).Msg("Login to ACR")
	command := NewCommand("az", ".")
	command.AddArgs(
		"acr",
		"login",
		"--subscription", subscription,
		"--name", acrName,
		"-u", clientID,
		"-p", clientSecret)
	return command.Run()
}
