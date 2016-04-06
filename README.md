# JIT - Jira & git [![Build Status](https://travis-ci.org/Rentlio/jit.svg?branch=master)](https://travis-ci.org/Rentlio/jit)

Jit is a simple terminal app that is used to help using git workflows with Jira Issue and Project tracking solution.



## Installation

Compile from source using go-lang compiler or [download](Downloads/) binary for your platform.


## Configuration
For jit to use your jira account you have to at least specify username, password and base URL. If you specify your jira project you will be able to use tool with only issues IDs without specifying project code.

There are 3 different ways you can configure your jit binary:

 * Using .jit.json file in your home folder
 * Specifying flags when using binary
 * Using environment variables

### Using .jit.json

Put .jit.json file in you home folder with content:

```json
{
  "username": "YOUR JIRA USERNAME",
  "password": "YOUR JIRA PASSWORD",
  "url" : "YOUR JIRA BASE URL",
  "project" : "OPTIONAL - JIRA PROJECT CODE",
  "working-branch" : "OPTIONAL - You project main git branch (master, develop,...)"
}
```

### Using cli flags

```sh
$ jit --username myuser --password mypass --url https://company.atlassian.net/rest/api/2/ --project PROJ ...
```

or using short codes for flags:

```sh
$ jit -u myuser -p mypass --url https://company.atlassian.net/rest/api/2/ --project PROJ ...
```

### Using environment variables
You just have to expose this environment variables:
 * JIRA_USERNAME
 * JIRA_PASSWORD
 * JIRA_URL
 * JIRA_PROJECT_CODE [optional]
 * JIT_WORKING_BRANCH [optional]

## Configuration options
  * username - required - Your jira username. Usually first part of email address you use for jira login
  * password - required - Your jira password
  * url - required - Url pointing on your jira rest api endpoint. Ussualy something like https://company.atlassian.net/rest/api/2/
  * project - optional - Your project code that is part of your issue IDs. In jira issues have IDs like PROJ-ID, if you specify project code as PROJ you will be able to use jit commands with only specifying number part of issue ID. This is useful if you only have one jira project and don't want to repeat this in each command.
  * working-branch - optional - If working branch is set, calling jit checkout without issue ID will checkout that branch.This is useful for switching to main working branch, for example develop, if using gitflow.

## Usage
To see full list of commands type:
```sh
$ jit -h
```

### Checkout existing branch

```sh
$ jit checkout 20
```

Or if you haven't specified Project code

```sh
$ jit checkout PROJ-20
```

Using short command
```sh
$ jit co 20
```

Checkouting working branch (if set in options)
```sh
$ jit co
```

### Creating new branch
Creating new branch is done same as checking out existing one with addition to -b flag. For example:
```sh
$ jit co -b 20
```
will create new branch with name set to issue ID 20 summary

### Listing issues for specific Version
To list all issues that are in some version you can call:
```sh
$ jit version v1.20.1
```
and list of issues will be shown. This is useful for deployment messages, notifications, etc...

## Contribution
This tool is actively used in Rentl.io Dev Workflow, but we are open to any changes, bugfixes, new features, etc. Just drop us some Pull Requests, making sure tests are passing, and we'll gladly merge it :).

## License
MIT
