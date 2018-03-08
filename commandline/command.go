package commandline

import (
	"os/exec"
)

// first argument is the command, like cat or echo,
// the second is the list of args to pass to it
type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

var runner Runner

// the real runner for the actual program, actually execs the command
func (r RealRunner) Run(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}

func Hello() string {
	out, err := runner.Run("echo", "hello")
	if err != nil {
		panic(err)
	}
	return string(out)
}
