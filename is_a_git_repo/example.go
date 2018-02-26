//usr/bin/env gorun $0;exit $?;
package main

import (
	"fmt"

	"github.com/ithaka/go-jira-pull-request/commandline"
)

func main() {
	pwd := commandline.Pwd()
	switch {
	case commandline.IsGitRepo():
		fmt.Println("Your local directory", pwd, "is a git repo")
	default:
		fmt.Println("Your local directory", pwd, "is NOT a git repo")
	}
}