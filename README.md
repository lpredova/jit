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
  "project" : "OPTIONAL - JIRA PROJECT CODE"
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

### Creating new branch
Creating new branch is done same as checking out existing one with addition to -b flag. For example:
```sh
$ jit co -b 20
```
will create new branch with name set to issue ID 20 summary

## Contribution
This tool is actively used in Rentl.io Dev Workflow, but we are open to any changes, bugfixes, new features, etc. Just drop us some Pull Requests, making sure tests are passing, and we'll gladly merge it :).

## License
MIT
