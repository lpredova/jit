package main

import (
	"testing"

	"github.com/codegangsta/cli"
)

// Testing that PasswordFlag will not show its value in help template
func TestPasswordFlagHelpString(t *testing.T) {
	flag := PasswordFlag{
		cli.StringFlag{},
		"password, p",
		"some value",
		"usage",
		"PASS_ENV",
		nil,
	}

	val := flag.String()
	if val != "--password, -p 	usage [$PASS_ENV]" {
		t.Error("--password, -p 	usage [$PASS_ENV], got:", val)
	}

}
