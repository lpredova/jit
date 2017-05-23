package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
)

var errorDecorator = ansi.ColorFunc("white:red")

// Create New Cli App
func createNewApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Jira & Git Worflow"
	app.Usage = "Simple tool for automating branch and projects management using jira issues"
	app.Author = "Rentl.io developers@rentl.io"
	app.Version = "0.5.0"

	return app
}

// Set global flags to App
func setGlobalFlags(app *cli.App, config *configuration) {
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "Username for jira basic auth",
			EnvVar:      "JIRA_USERNAME",
			Value:       config.Username,
			Destination: &config.Username,
		},
		cli.StringFlag{
			Name:        "password, p",
			Usage:       "Password for jira Basic Auth",
			EnvVar:      "JIRA_PASSWORD",
			Value:       config.Password,
			Destination: &config.Password,
		},
		cli.StringFlag{
			Name:        "url",
			Usage:       "Base url for your jira api",
			EnvVar:      "JIRA_URL",
			Value:       config.URL,
			Destination: &config.URL,
		},
		cli.StringFlag{
			Name:  "projects, pro",
			Usage: "Lists jira projects in configuration",
		},
		cli.StringFlag{
			Name:   "alias, al",
			Usage:  "Jira project alias, User defined alias used for easier project managment.",
			EnvVar: "JIRA_PROJECT_CODE",
		},
		cli.StringFlag{
			Name:   "working-branch, wb",
			Usage:  "Git working branch. If set, checkout command without ID will checkout this branch.",
			EnvVar: "JIT_WORKING_BRANCH",
		},
	}
}

// Set application commands
func setCommands(app *cli.App, config *configuration) {
	app.Commands = []cli.Command{
		{
			Name:    "checkout",
			Aliases: []string{"co"},
			Usage:   "Checkout branch",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "b",
					Usage: "Create new branch while checkout",
				},
			},
			Action: func(c *cli.Context) {
				valid := validateConfiguration(config)
				if !valid {
					fmt.Println(errorDecorator("Please provide valid configuration"))
					os.Exit(1)
				}
				checkoutBranch(c, config)
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Get version issues",
			Action: func(c *cli.Context) {
				valid := validateConfiguration(config)
				if !valid {
					fmt.Println(errorDecorator("Please provide valid configuration"))
					os.Exit(1)
				}
				getVersionIssues(c, config)
			},
		},
		{
			Name:    "description",
			Aliases: []string{"d"},
			Usage:   "Show issue description",
			Action: func(c *cli.Context) {
				valid := validateConfiguration(config)
				if !valid {
					fmt.Println(errorDecorator("Please provide valid configuration"))
					os.Exit(1)
				}
				showIssueDetails(c, config)
			},
		},
		{
			Name:    "projects",
			Aliases: []string{"pro"},
			Usage:   "List configured jira projects",
			Action: func(c *cli.Context) {
				valid := validateConfiguration(config)
				if !valid {
					fmt.Println(errorDecorator("Please provide valid configuration"))
					os.Exit(1)
				}

				listJiraProjectsFromConfiguration(c, config)
			},
		},
	}
}

func main() {
	config := getJSONConfiguration()
	app := createNewApp()

	setGlobalFlags(app, &config)
	setCommands(app, &config)
	app.Run(os.Args)
}
