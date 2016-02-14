package main

import (
	"encoding/json"
	"testing"
)

func TestStripProjectCode(t *testing.T) {

	val := stripProjectCode("PROJ-123")
	if val != "123" {
		t.Error("Expected 123, got:", val)
	}

	val = stripProjectCode("456")
	if val != "456" {
		t.Error("Expected 456, got:", val)
	}

	val = stripProjectCode("")
	if val != "" {
		t.Error("Expected empty string, got:", val)
	}

}

func TestGetBranchNameForIssue(t *testing.T) {
	var issue Issue
	jsonContent := `{"key":"PROJ-123", "fields":{"summary":"some issue summary","issuetype":{"name":"Story"}}}`
	json.Unmarshal([]byte(jsonContent), &issue)
	branchName, err := getBranchNameForIssue(issue)
	if err != nil || branchName != "story-123-some-issue-summary" {
		t.Error(`Expected branch name:"story-123-some-issue-summary", got:`, branchName)
	}

	jsonContent = `{"key":"PROJ-123", "fields":{"summary":"some issue #special summary","issuetype":{"name":"Story"}}}`
	json.Unmarshal([]byte(jsonContent), &issue)
	branchName, err = getBranchNameForIssue(issue)
	if err != nil || branchName != "story-123-some-issue-special-summary" {
		t.Error(`Expected branch name:"story-123-some-issue-special-summary", got:`, branchName)
	}

}

func TestGetJiraIssuesRestURL(t *testing.T) {
	// Testing URL when Project Code is specified in configuration
	config := configuration{
		"username",
		"pass",
		"https://project.atlassian.net/rest/api/2/",
		"PROJ",
	}
	url := getJiraIssuesRestURL("123", &config)
	expected := "https://project.atlassian.net/rest/api/2/issue/PROJ-123"
	if url != expected {
		t.Error("Expected url:", expected, ", got:", url)
	}

	// Testing URL when project code is part of ID
	config = configuration{
		"username",
		"pass",
		"https://project.atlassian.net/rest/api/2/",
		"",
	}
	url = getJiraIssuesRestURL("NEW-123", &config)
	expected = "https://project.atlassian.net/rest/api/2/issue/NEW-123"
	if url != expected {
		t.Error("Expected url:", expected, ", got:", url)
	}
}
