package main

import "testing"

func TestGetJSONConfoguration(t *testing.T) {
	getJSONConfiguration()
}

func TestValidateConfigurationFailed(t *testing.T) {
	config := configuration{}
	if isValid := validateConfiguration(&config); isValid == true {
		t.Error("Expected false, got ", isValid)
	}
}

func TestValidateConfigurationSuccess(t *testing.T) {
	config := configuration{
		Username: "user",
		Password: "pass",
		URL:      "https://test.jira.com/api",
	}
	if isValid := validateConfiguration(&config); isValid == false {
		t.Error("Expected true, got ", isValid)
	}
}

func TestLoadUnexsitentConfigFile(t *testing.T) {
	_, err := loadConfigFile()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
}
