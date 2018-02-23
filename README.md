# go-jira-pull-request
Utility that creates a github pull request with github markdown translated from JIRA story.

## Setup

Install Go.

If desired, add the following lines to your bash_profile:
```bash
# Work with and use Go
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)
```

To clone this repo, you can simply use `go get github.com/ithaka/go-jira-pull-request`. Doing so initializes the repo in your `$GOPATH`, which makes a number of things about development significantly easier.

You can run `cd $GOPATH/src/github.com/ithaka/go-jira-pull-request/is_a_git_repo ; go install` to install the command `is_a_git_repo`. If you've added the the above lines to your bash profile, you should be able to run that from anywhere and get a very simple output about if the directory you run it from is or is inside of a git repo


## Tasks

These are intended to be independent units of work that could be done in any order.
The intention is that pairs of developers could work on a particular task. At some point they could be combined into a complete solution.

If you can execute programs from GoLang, you can cheat and use any piece(s) of the reference bash/javascript/ruby solution. You can also use any GoLang libraries if they are convenient. Or not. Whatever!

### Exercise 0: Execute a unix utility from GoLang

Anything, really!

### Exercise 1: Make a web request from JIRA

Equivalent of:
```
curl -u "$JIRA_LOGIN:$JIRA_PASSWORD" \
-X GET \
--write '\n%{http_code}\n' \
--fail \
--silent \
"${JIRA_URL}${JIRA_API_URI}issue/${1}"
```

### Exercise 2: Convert jira markdown to github markdown with regular expressions

[J2M](https://github.com/asharpe/J2M/blob/all-fixes/src/J2M.js#L12) has a lot of useful regexes. Maybe just numbered lists and unordered lists to begin with.

### Exercise 3: Determine git branch of current working directory

Equivalent of:
```
git rev-parse --abbrev-ref HEAD
```

### Exercise 4: Determine github repo
Equivalent of:
```
# try the upstream branch if possible, otherwise origin will do
upstream=$(git config --get remote.upstream.url)
origin=$(git config --get remote.origin.url)
if [ -z $upstream ]; then
  upstream=$origin
fi

to_user=$(echo $upstream | sed -e 's/.*[\/:]\([^/]*\)\/[^/]*$/\1/')
from_user=$(echo $origin | sed -e 's/.*[\/:]\([^/]*\)\/[^/]*$/\1/')
repo=$(basename `git rev-parse --show-toplevel`)
```

### Exercise 5: Make Pull Request

Equivalent of:
```
hub pull-request -o -m "${MESSAGE}"
```
Maybe try [Go Github](https://github.com/google/go-github) or [githubql](https://github.com/shurcooL/githubql)
