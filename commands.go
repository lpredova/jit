package main

import (
	"fmt"
	"os/exec"

	"github.com/codegangsta/cli"
)

// Use Jira issue ID to form a branch name
// and checkout to that branch. If -b switch is specified
// new branch will be created
func checkoutBranch(c *cli.Context, config *configuration) {
	branchName, err := getBranchName(c.Args().First(), config)
	if err != nil {
		fmt.Println(err.Error())
	}
	create := c.Bool("b")
	var out []byte
	if create {
		out, err = exec.Command("git", "checkout", "-b", branchName).CombinedOutput()
	} else {
		out, err = exec.Command("git", "checkout", branchName).CombinedOutput()
	}
	if err != nil {
		fmt.Printf("%s", errorDecorator(string(out[:])))
	} else {
		fmt.Printf("%s", out)
	}
}
