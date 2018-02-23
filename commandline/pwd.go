// Package commandline contains utilities for interacting directly with command line utilities
package commandline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Man, I go no unit test!
func Pwd() string {
	if cmdOut, err := exec.Command("pwd").Output() ; err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running the command: ", err)
		return ""
	} else {
	    return strings.TrimSpace(string(cmdOut))
	}
}