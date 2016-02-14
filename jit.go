package main

import (
	"os"

	"github.com/codegangsta/cli"
)

// Create New Cli App
func createNewApp() *cli.App {
	app := cli.NewApp()

	app.Name = "Jira & Git Worflow"
	app.Usage = "Simple tool for automating branch management using jira issues"
	app.Author = "Rentl.io developers@rentl.io"
	app.Version = "0.1.0"

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
		PasswordFlag{
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
			Name:        "project",
			Usage:       "Jira project code. If specified only issues ID can be used in commands.",
			EnvVar:      "JIRA_PROJECT_CODE",
			Value:       config.ProjectCode,
			Destination: &config.ProjectCode,
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
				checkoutBranch(c, config)
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
