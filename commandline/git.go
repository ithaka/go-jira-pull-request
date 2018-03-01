// Package commandline contains utilities for interacting directly with command line utilities
package commandline

import "os/exec"

// Man, I go no unit test!
func IsGitRepo() bool {
	cmd := exec.Command
	return IsGitRepoWithExec(cmd)
}


func IsGitRepoWithExec(cmd func(string, ...string) *exec.Cmd ) bool {
	if _, err := cmd("git", "rev-parse", "--is-inside-work-tree").Output() ; err != nil {
		return false
	} else {
		return true
	}

}
