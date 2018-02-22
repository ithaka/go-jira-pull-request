package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	// I would prefer to use commandline.Pwd(), but my import doesn't seem to work...
	pwd := Pwd()
	switch IsGitRepo() {
	case true:
		fmt.Println(pwd, " is a git repo")
	default:
		fmt.Println(pwd, " is NOT a git repo")
	}
}

func Pwd() string {
	if cmdOut, err := exec.Command("pwd").Output() ; err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running the command: ", err)
		return ""
	} else {
		return string(cmdOut)
	}
}

func IsGitRepo() bool {
	if _, err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output() ; err != nil {
		return false
	} else {
		return true
	}
}