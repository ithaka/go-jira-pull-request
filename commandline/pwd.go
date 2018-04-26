// Package commandline contains utilities for interacting directly with command line utilities
package commandline

import (
	"os/exec"
	"strings"
)

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

var runner Runner = RealRunner{}

func (r RealRunner) Run(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}
// Man, I go no unit test!
func Pwd() (string, error) {
	out, err := runner.Run("pwd")
	return strings.TrimSpace(string(out)), err
}