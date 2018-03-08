package main

import (
	"fmt"
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
// Untestable via Unit test
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

// Actually executes the real calls as well, therefor also untestable
func main() {
	// This dependency injection works because it's in the same package
	runner = RealRunner{}
	fmt.Println(Hello())
	// This does NOT work because I can't find a good way to set commandline.runner = RealRunner{}
	//fmt.Println(commandline.Hello())
}