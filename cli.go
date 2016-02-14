package main

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/mgutz/ansi"
)

var errorDecorator = ansi.ColorFunc("white:red")

// PasswordFlag - Override Flag String Method for not showing values of flags
// So username and password will not be shown in cli
type PasswordFlag struct {
	cli.StringFlag
	Name        string
	Value       string
	Usage       string
	EnvVar      string
	Destination *string
}

// String returns the usage
func (f PasswordFlag) String() string {
	var fmtString string
	fmtString = "%s \t%v"
	return withEnvHint(f.EnvVar, fmt.Sprintf(fmtString, prefixedNames(f.Name), f.Usage))
}

func withEnvHint(envVar, str string) string {
	envText := ""
	if envVar != "" {
		envText = fmt.Sprintf(" [$%s]", strings.Join(strings.Split(envVar, ","), ", $"))
	}
	return str + envText
}

func prefixedNames(fullName string) (prefixed string) {
	parts := strings.Split(fullName, ",")
	for i, name := range parts {
		name = strings.Trim(name, " ")
		prefixed += prefixFor(name) + name
		if i < len(parts)-1 {
			prefixed += ", "
		}
	}
	return
}

func prefixFor(name string) (prefix string) {
	if len(name) == 1 {
		prefix = "-"
	} else {
		prefix = "--"
	}

	return
}
