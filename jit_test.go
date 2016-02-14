package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCreateNewApp(t *testing.T) {
	app := createNewApp()
	if app.Name != "Jira & Git Worflow" {
		t.Error("Expected app name is 'Jira & Git Worflow', got:", app.Name)
	}
}

func TestSetGlobalFlags(t *testing.T) {
	app := createNewApp()
	config := configuration{}
	setGlobalFlags(app, &config)
	if numFlags := len(app.Flags); numFlags == 0 {
		t.Error("Expected App flags to be defined")
	}
}

func TestSetCommands(t *testing.T) {
	app := createNewApp()
	config := configuration{}
	setCommands(app, &config)
	if numCommands := len(app.Commands); numCommands == 0 {
		t.Error("Expected App Commands to be defined")
	}
}

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out[:]), "Jira & Git Worflow") {
		t.Error("Help output expected")
	}

}
