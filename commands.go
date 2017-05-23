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
// second param is project alias used to switch between projects (optional)
// if second param is not provided then we use project from config with default flag
func checkoutBranch(c *cli.Context, config *configuration) {

	issueID := c.Args().First()

	projectAlias := c.String("pr")
	if len(projectAlias) == 0 {
		projectAlias = getDefaultProjectAlias(config)
	}

	branchName, err := getBranchName(issueID, projectAlias, config)
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

func showIssueDetails(c *cli.Context, config *configuration) {
	issue, err := GetIssue(c.Args().First(), "", config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(issue)
}

func listJiraProjectsFromConfiguration(c *cli.Context, config *configuration) {
	if len(config.Projects) > 0 {
		for index, project := range config.Projects {
			project := fmt.Sprintf("%d.\nCode:\t%s\nAlias:\t%s\nBranch:\t%s\n", index+1, project.ProjectCode, project.Alias, project.WorkingBranch)
			fmt.Println(project)
		}
	} else {
		fmt.Println("No configured projects")
	}
}
