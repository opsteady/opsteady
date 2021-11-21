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
}

// NewTerraform returns a Terraform task runner
func NewTerraform(path, backendConfigPath, backendStorageName, cachePath string, logger *zerolog.Logger) *Terraform {
	return &Terraform{
		path:               path,
		backendConfigPath:  backendConfigPath,
		backendStorageName: backendStorageName,
		cachePath:          fmt.Sprintf("%s/terraform", cachePath),
		logger:             logger,
	}
}

// FmtCheck checks the formatting of Terraform code
func (t *Terraform) FmtCheck() error {
	t.logger.Info().Str("path", t.path).Msg("Checking Terraform code format in path")

	command := NewCommand("terraform", t.path)
	command.AddArgs("fmt", "-check", ".")

	return command.Run()
}

// Plan runs a Terraform plan.
// The vars are added as a JSON file so that Terraform runs without input questions.
func (t *Terraform) Plan(varsFile string) error {
	t.logger.Info().Str("path", t.path).Msg("Running Terraform plan in")

	command := NewCommand("terraform", t.path)
	command.AddArgs("plan", "-compact-warnings", "-input=false", fmt.Sprintf("-var-file=%s", varsFile))

	return command.Run()
}

// Apply runs a Terraform apply with auto approval.
// The vars are added as TF_VAR environment variables when executing,
// so that Terraform runs without input questions.
func (t *Terraform) Apply(varsFile string) error {
	t.logger.Info().Str("path", t.path).Msg("Running Terraform apply in")

	command := NewCommand("terraform", t.path)
	command.AddArgs("apply", "-compact-warnings", "-input=false", "-auto-approve", fmt.Sprintf("-var-file=%s", varsFile))

	return command.Run()
}

// InitAndPlan combines Terraform init and plan
func (t *Terraform) InitAndPlan(varsFile string) error {
	if err := t.Init(); err != nil {
		return err
	}
	return t.Plan(varsFile)
}

// InitAndApply combines Terraform init and apply
func (t *Terraform) InitAndApply(varsFile string) error {
	if err := t.Init(); err != nil {
		return err
	}
	return t.Apply(varsFile)
}

// Init initializes the Terraform providers and backend.
// Terraform is run with reconfigure and upgrade flags.
// It sets the TF_PLUGIN_CACHE_DIR to cache plugins in a well-known location.
func (t *Terraform) Init() error {
	t.logger.Info().Str("path", t.path).Str("cache", t.cachePath).Msg("Initializing Terraform in")

	// Always remove the .terraform folder to make sure that we use the correct platform-dependent providers.
	if err := os.RemoveAll(fmt.Sprintf("%s/.terraform", t.path)); err != nil {
		return err
	}

	command := NewCommand("terraform", t.path)
	command.AddArgs("init", "-reconfigure", "-upgrade")
	if t.backendStorageName != "" {
		command.AddArgs(fmt.Sprintf("-backend-config=storage_account_name=%s", t.backendStorageName))
	}
	if t.backendConfigPath != "" {
		command.AddArgs(fmt.Sprintf("-backend-config=%s", t.backendConfigPath))
	}

	command.AddEnv("TF_PLUGIN_CACHE_DIR", t.cachePath)

	return command.Run()
}

// Destroy runs a Terraform destroy command
func (t *Terraform) Destroy(varsFile string) error {
	if err := t.Init(); err != nil {
		return err
	}

	t.logger.Info().Str("path", t.path).Msg("Running Terraform destroy in")

	command := NewCommand("terraform", t.path)
	command.AddArgs("destroy", "-compact-warnings", "-auto-approve", fmt.Sprintf("-var-file=%s", varsFile))

	return command.Run()
}
