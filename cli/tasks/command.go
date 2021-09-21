package tasks

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

// Command is the interface for running binaries
type Command interface {
	Run() (string, error)
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
	writer     io.Writer
}

// NewCommand creates the CommandImpl
func NewCommand(command, workingDir string, writer io.Writer) Command {
	return &CommandImpl{
		Args:       []string{},
		Env:        make(map[string]string),
		Command:    command,
		WorkingDir: workingDir,
		writer:     writer,
	}
}

// Run runs the command
func (c *CommandImpl) Run() (string, error) {
	cmd := exec.Command(c.Command, c.Args...)
	cmd.Dir = c.WorkingDir

	var stdoutBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(c.writer, &stdoutBuf)
	cmd.Stderr = c.writer

	// add all environment variables to the Env of the command
	env := os.Environ()
	for key, value := range c.Env {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env

	if err := cmd.Run(); err != nil {
		return "", errors.Wrapf(err, "Failed to execute command %s", c.Command)
	}

	return stdoutBuf.String(), nil
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
