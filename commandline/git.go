// Package commandline contains utilities for interacting directly with command line utilities
package commandline

import "os/exec"

// Man, I go no unit test!
func IsGitRepo() bool {
	if _, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output() ; err != nil {
		return false
	} else {
		return true
	}
}
