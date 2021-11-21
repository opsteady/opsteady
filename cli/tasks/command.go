package tasks

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// Command is the interface for running binaries
type Command interface {
	Run() error
	AddArgs(...string)
	AddEnv(string, string)
	AddEnvs(map[string]string)
}

// CommandImpl is the implementation of the Command interface
type CommandImpl struct {
	Command    string            // The command to run
	Args       []string          // The args to pass to the command
	WorkingDir string            // The working directory
	Env        map[string]string // Additional environment variables to set
}

// NewCommand creates the CommandImpl
func NewCommand(command, workingDir string) Command {
	return &CommandImpl{
		Args:       []string{},
		Env:        make(map[string]string),
		Command:    command,
		WorkingDir: workingDir,
	}
}

// Run runs the command
func (c *CommandImpl) Run() error {
	cmd := exec.Command(c.Command, c.Args...)
	cmd.Dir = c.WorkingDir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// add all environment variables to the Env of the command
	env := os.Environ()
	for key, value := range c.Env {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "Failed to execute command %s", c.Command)
	}

	return nil
}

// AddArgs adds arguments to the command
func (c *CommandImpl) AddArgs(args ...string) {
	c.Args = append(c.Args, args...)
}

// AddEnv adds value as environment variable of the command
func (c *CommandImpl) AddEnv(key, value string) {
	c.Env[key] = value
}

// AddEnvs adds values as environment variables of the command
func (c *CommandImpl) AddEnvs(values map[string]string) {
	for k, v := range values {
		c.Env[k] = v
	}
}
