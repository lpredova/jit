package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Issue - Jira issue data struct
type Issue struct {
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
	Fields struct {
		Summary     string `json:"summary,omitempty"`
		Description string `json:"description,omitempty"`
		IssueType   struct {
			ID          string `json:"id,omitempty"`
			Self        string `json:"self,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
			Subtask     bool   `json:"subtask,omitempty"`
			IconURL     string `json:"iconUrl,omitempty"`
		} `json:"issuetype,omitempty"`
		Parent      *Issue           `json:"parent,omitempty"`
		Subtasks    []*Issue         `json:"subtasks,omitempty"`
		Assignee    *User            `json:"assignee,omitempty"`
		FixVersions []*Version       `json:"fixVersions,omitempty"`
		Labels      []string         `json:"labels,omitempty"`
		Status      *IssueStatus     `json:"status,omitempty"`
		Resolution  *IssueResolution `json:"resolution,omitempty"`
	} `json:"fields,omitempty"`
}

// User - jira user struct
type User struct {
	Self         string      `json:"self,omitempty"`
	Key          string      `json:"key,omitempty"`
	Name         string      `json:"name,omitempty"`
	EmailAddress string      `json:"emailAddress,omitempty"`
	AvatarURLs   *AvatarURLs `json:"avatarUrls,omitempty"`
	DisplayName  string      `json:"displayName,omitempty"`
	Active       bool        `json:"active,omitempty"`
	TimeZone     string      `json:"timeZone,omitempty"`
	Groups       struct {
		Size  int `json:"size,omitempty"`
		Items []struct {
			Name string `json:"name,omitempty"`
			Self string `json:"self,omitempty"`
		} `json:"items,omitempty"`
	} `json:"groups,omitempty"`
}

// AvatarURLs - jira avatar struct
type AvatarURLs struct {
	Size16 string `json:"16x16,omitempty"`
	Size24 string `json:"24x24,omitempty"`
	Size32 string `json:"32x32,omitempty"`
	Size48 string `json:"48x48,omitempty"`
}

// Version - Jira version struct
type Version struct {
	ID            string `json:"id,omitempty"`
	Self          string `json:"self,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Project       string `json:"project,omitempty"`
	ProjectID     int    `json:"projectId,omitempty"`
	Released      bool   `json:"released,omitempty"`
	Archived      bool   `json:"archived,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	UserStartDate string `json:"userStartDate,omitempty"`
}

// IssueStatus - jira issue status struct
type IssueStatus struct {
	ID          string `json:"id,omitempty"`
	Self        string `json:"self,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
}

// IssueResolution - jira issue resolution struct
type IssueResolution struct {
	ID          string `json:"id,omitempty"`
	Self        string `json:"self,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// VersionIssues - struct for holding issues fixed in Version
type VersionIssues struct {
	Total  int      `json:"total,omitempty"`
	Issues []*Issue `json:"issues,omitempty"`
}

func getJiraVersionIssues(version string, config *configuration) (VersionIssues, error) {
	client := &http.Client{}
	var issues VersionIssues
	url := strings.TrimRight(config.URL, "/") + "/search"

	var data = []byte(`{"jql":"fixVersion = ` + version + `","startAt":0,"maxResults":1000,"fields":["id","key","summary", "issuetype"],"expand":[]}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic("Error while building jira request")
	}
	req.SetBasicAuth(config.Username, config.Password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return issues, err
	}

	error := json.Unmarshal(contents, &issues)
	return issues, error
}

// GetIssue - Get Jira issue data
func GetIssue(id string, config *configuration) (Issue, error) {
	client := &http.Client{}
	var issue Issue

	url := getJiraIssuesRestURL(id, config)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Error while building jira request")
	}
	req.SetBasicAuth(config.Username, config.Password)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return issue, err
	}

	error := json.Unmarshal(contents, &issue)
	return issue, error
}

func getJiraIssuesRestURL(id string, config *configuration) string {
	var url string
	baseURL := strings.TrimRight(config.URL, "/")
	if config.ProjectCode != "" {
		url = baseURL + "/issue/" + config.ProjectCode + "-" + id
	} else {
		url = baseURL + "/issue/" + id
	}
	return url
}

// GetBranchNameForIssue - Return branch name jira issue
func getBranchNameForIssue(issue Issue) (string, error) {
	regex, _ := regexp.Compile("[^A-Za-z0-9- ]+")
	summary := strings.Replace(strings.ToLower(regex.ReplaceAllString(issue.Fields.Summary, "")), " ", "-", -1)
	if summary == "" {
		return "", errors.New("Unable to find issue with id:" + issue.Key)
	}
	issueType := strings.ToLower(issue.Fields.IssueType.Name)
	return issueType + "-" + stripProjectCode(issue.Key) + "-" + summary, nil
}

func getBranchName(id string, config *configuration) (string, error) {
	issue, err := GetIssue(id, config)
	if err != nil {
		return "", err
	}
	return getBranchNameForIssue(issue)
}

func stripProjectCode(id string) string {
	parts := strings.Split(id, "-")
	return parts[len(parts)-1]
}
