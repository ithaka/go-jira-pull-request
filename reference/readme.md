## Create Github Pull Request from commmand line using git and JIRA information

This script will use the git topic branch name to look up the JIRA story.
It will pull the summary and description from JIRA assuming the topic branch name matches the JIRA issue.
It converts the Jira flavored markdown to github flavored markdown.
It will create a github pull request with github flavored markdown description. 

__NOTE:__ Change the `JIRA_URL` to match your own!

### Installation
```
$ npm install -g j2m
$ brew install jq curl hub
$ source git_jira_github_pull_request.sh
```

### Usage
```
$ make_pull
```

### Caveats
 * Adding `JIRA_LOGIN` and/or `JIRA_PASSWORD` to the environment will mean you won't be prompted for them.
 * If you have local changes that have not been pushed to origin, you get a horrible error message `Unprocessable Entity (HTTP 422) Invalid value for "head"`. Just push first and you're ok.
 * Hub requires an access token, and if you don't have one, on first invocation it will ask for your github userid and password to generate one for you. You can avoid this by doing it manually:
   1. [Create a Personal Access Token on Github.](https://github.com/settings/tokens) This acts as an OAuth token
   2. Save the following information (substitute `ACCESS_TOKEN` with the real value) in ~/.config/hub:
```
github.com:
- user: StevenACoffman
  oauth_token: ACCESS_TOKEN
  protocol: https
```
 * If you don't have j2m installed, it falls back to just using the Jira markdown