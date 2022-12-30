package tasks

import (
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
