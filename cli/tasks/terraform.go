package tasks

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Terraform represents a Terraform task library
type Terraform struct {
	path               string
	backendConfigPath  string
	cachePath          string
	backendStorageName string
	logger             *zerolog.Logger
	logWriter          *LogWriter
}

// NewTerraform returns a Terraform task runner
func NewTerraform(path, backendConfigPath, backendStorageName, cachePath string, logger *zerolog.Logger) *Terraform {
	return &Terraform{
		path:               path,
		backendConfigPath:  backendConfigPath,
		backendStorageName: backendStorageName,
		cachePath:          fmt.Sprintf("%s/terraform", cachePath),
		logger:             logger,
		logWriter:          &LogWriter{logger: logger},
	}
}

// FmtCheck checks the formatting of Terraform code
func (t *Terraform) FmtCheck() error {
	t.logger.Info().Str("path", t.path).Msg("Checking Terraform code format in path")

	command := NewCommand("terraform", t.path, t.logWriter)
	command.AddArgs("fmt", "-check", ".")

	_, err := command.Run()
	return err // either it is nil or it has an error
}

// Plan runs a Terraform plan
// The vars are added as TF_VAR environment variables when executing
// Terraform is run without input question
func (t *Terraform) Plan(vars map[string]string) error {
	t.logger.Info().Str("path", t.path).Msg("Running Terraform plan in")

	command := NewCommand("terraform", t.path, t.logWriter)
	command.AddArgs("plan", "-input=false")
	command.AddEnvs(addTfVarPrefixToVars(vars))

	_, err := command.Run()
	return err // either it is nil or it has an error
}

// Apply runs a Terraform apply
// The vars are added as TF_VAR environment variables when executing
// Terraform is run without input question and with auto approve
func (t *Terraform) Apply(vars map[string]string) error {
	t.logger.Info().Str("path", t.path).Msg("Running Terraform apply in")

	command := NewCommand("terraform", t.path, t.logWriter)
	command.AddArgs("apply", "-input=false", "-auto-approve")
	command.AddEnvs(addTfVarPrefixToVars(vars))

	_, err := command.Run()
	return err // either it is nil or it has an error
}

// InitAndPlan combines Terraform init and plan
func (t *Terraform) InitAndPlan(vars map[string]string) error {
	if err := t.Init(); err != nil {
		return err
	}
	return t.Plan(vars)
}

// InitAndApply combines Terraform init and apply
func (t *Terraform) InitAndApply(vars map[string]string) error {
	if err := t.Init(); err != nil {
		return err
	}
	return t.Apply(vars)
}

// Init initializes the Terraform providers and backend
// Always remove the .terraform folder to make sure that we use the correct platform-dependent providers.
// Terraform is run with reconfigure and upgrade flags
// It sets the TF_PLUGIN_CACHE_DIR
func (t *Terraform) Init() error {
	t.logger.Info().Str("path", t.path).Str("cache", t.cachePath).Msg("Initializing Terraform in")

	if err := os.RemoveAll(fmt.Sprintf("%s/.terraform", t.path)); err != nil {
		return err
	}

	command := NewCommand("terraform", t.path, t.logWriter)
	command.AddArgs("init", "-reconfigure", "-upgrade")
	if t.backendStorageName != "" {
		command.AddArgs(fmt.Sprintf("-backend-config=storage_account_name=%s", t.backendStorageName))
	}
	if t.backendConfigPath != "" {
		command.AddArgs(fmt.Sprintf("-backend-config=%s", t.backendConfigPath))
	}
	command.AddEnv("TF_PLUGIN_CACHE_DIR", t.cachePath)

	_, err := command.Run()
	return err // either it is nil or it has an error
}

func addTfVarPrefixToVars(vars map[string]string) map[string]string {
	tmp := make(map[string]string)
	for k, v := range vars {
		tmp[fmt.Sprintf("TF_VAR_%s", k)] = v
	}
	return tmp
}

// Destroy runs a Terraform destroy command
func (t *Terraform) Destroy(vars map[string]string) error {
	if err := t.Init(); err != nil {
		return err
	}

	t.logger.Info().Str("path", t.path).Msg("Running Terraform destroy in")

	command := NewCommand("terraform", t.path, t.logWriter)
	command.AddArgs("destroy", "-auto-approve")
	command.AddEnvs(addTfVarPrefixToVars(vars))

	_, err := command.Run()
	return err // either it is nil or it has an error
}
