package tasks

import (
	"os"
	"testing"
)

func TestCommandImpl_Run(t *testing.T) {
	// create a command that prints "Hello, World!"
	cmd := NewCommand("echo", "")
	cmd.AddArgs("Hello, World!")

	// run the command and check the output
	if err := cmd.Run(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCommandImpl_AddArgs(t *testing.T) {
	// create a command that prints "Hello, World!"
	cmd := NewCommand("echo", "").(*CommandImpl)
	cmd.AddArgs("Hello, World!")

	// check that the correct arguments were added
	expectedArgs := []string{"Hello, World!"}
	if !equal(cmd.Args, expectedArgs) {
		t.Errorf("incorrect args. expected: %v, got: %v", expectedArgs, cmd.Args)
	}
}

func TestCommandImpl_AddTarget(t *testing.T) {
	// create a command that sets an environment variable
	cmd := NewCommand("env", "")
	cmd.AddTarget("MY_VAR", "foo")

	// run the command and check the output
	if err := cmd.Run(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check that the environment variable was set correctly
	if os.Getenv("MY_VAR") != "foo" {
		t.Errorf("incorrect value for MY_VAR. expected: foo, got: %s", os.Getenv("MY_VAR"))
	}
}

func TestCommandImpl_AddTargets(t *testing.T) {
	// create a command that sets multiple environment variables
	cmd := NewCommand("env", "")
	cmd.AddTargets(map[string]string{
		"MY_VAR_1": "foo",
		"MY_VAR_2": "bar",
	})

	// run the command and check the output
	if err := cmd.Run(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check that the environment variables were set correctly
	if os.Getenv("MY_VAR_1") != "foo" {
		t.Errorf("incorrect value for MY_VAR_1. expected: foo, got: %s", os.Getenv("MY_VAR_1"))
	}
	if os.Getenv("MY_VAR_2") != "bar" {
		t.Errorf("incorrect value for MY_VAR_2. expected: bar, got: %s", os.Getenv("MY_VAR_2"))
	}
}

// equal is a helper function that checks if two slices are equal
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
