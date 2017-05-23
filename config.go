package main

import (
	"encoding/json"
	"os"
	"os/user"
)

// configuration struct for holding config data
type configuration struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	URL      string    `json:"url"`
	Projects []Project `json:"projects"`
}

// Project is config for each project in jira separately
type Project struct {
	Alias         string `json:"alias"`
	ProjectCode   string `json:"project"`
	WorkingBranch string `json:"working-branch"`
}

//const jitCoinfigFIle = ".jit.json"
const jitCoinfigFIle = ".jit2.json"

// Try to load configuration from json file.
// If unable to load return empty config struct
func getJSONConfiguration() configuration {
	var configuration = configuration{}

	file, err := loadConfigFile(jitCoinfigFIle)
	if err != nil {
		return configuration
	}
	decoder := json.NewDecoder(file)

	decoder.Decode(&configuration)
	return configuration
}

// Try to load config json file, in user home folder
func loadConfigFile(fileName string) (*os.File, error) {
	file := &os.File{}
	usr, err := user.Current()
	if err != nil {
		return file, err
	}
	file, err = os.Open(usr.HomeDir + "/" + fileName)
	if err != nil {
		return file, err
	}
	return file, nil
}

// Validate loaded configuration
// username, password and url are required values for jit to work
func validateConfiguration(conf *configuration) bool {
	if conf.Username == "" || conf.Password == "" || conf.URL == "" {
		return false
	}

	return true
}
