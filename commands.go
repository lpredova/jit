package main

import (
	"fmt"
	"os"
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
		os.Exit(1)
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

func getVersionIssues(c *cli.Context, config *configuration) {
	issues, err := getJiraVersionIssues(c.Args().First(), config)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(issues.Issues) == 0 {
		fmt.Println(errorDecorator("No issues found in version:" + c.Args().First()))
		os.Exit(1)
	}
	for _, issue := range issues.Issues {
		fmt.Println(issue.Fields.IssueType.Name, "-", issue.Fields.Summary)
	}
}
