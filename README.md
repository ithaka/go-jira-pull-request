# go-jira-pull-request
Utility that creates a github pull request with github markdown translated from JIRA story.


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

JIRA Issues uses [Wiki Markup](https://jira.atlassian.com/secure/WikiRendererHelpAction.jspa?section=all). Github Pull Request descriptions use [Github Markdown](https://guides.github.com/features/mastering-markdown/). Converting between them can be done in a number of ways.

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
### Exercise 5: Generate Github Personal Access Token

Note: Use Basic Auth once to [create an OAuth2 token](http://developer.github.com/v3/oauth/#oauth-authorizations-api)

```
read -s MFA_OTP; export MFA_OTP
curl https://api.github.com/authorizations \
--user "caspyin" \
-H "X-GitHub-OTP: $MFA_OTP" \
--data '{"scopes":["repo"],"note":"Demo"}'
```
This will prompt you for your GitHub Multi-Factor One Time Password, then your regular password, and return your OAuth token in the response. It will also create a new Authorized application [in your account settings](https://github.com/settings/applications).


### Exercise 6: Make Pull Request

Equivalent of:
```
hub pull-request -o -m "${MESSAGE}"
```
GitHub's API requires authentication, the simplest way is to use a Personal Access Token, and setting the environment's GITHUB_TOKEN to this value.

```
export GO_JIRA_PULL_REQUEST_AUTH_TOKEN=aabbcc...ddeeff
```
make_pull then takes the GitHub's user's login as the first and only argument, and it's required.

```
make_pull mynameisawesome
```

#### Personal Access tokens

Check them out here https://github.com/settings/tokens


#### OAuth 

The first thing to know is that your API Token (found in https://github.com/settings/admin) is not the same token used by OAuth. They are different tokens and you will need to generate an OAuth token to be authorized.

Follow the API's instructions at http://developer.github.com/v3/oauth/ under the sections "Non-Web Application Flow" and "Create a new authorization" to become authorized.

Now that you have the OAuth token there are two ways to use the token to make requests that require authentication (replace "OAUTH-TOKEN" with your actual token)

    curl https://api.github.com/gists/starred?access_token=OAUTH-TOKEN
    curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com/gists/starred

List the authorizations you already have

    curl --user "caspyin" https://api.github.com/authorizations

#### Make a pull request

 https://developer.github.com/v3/pulls/#create-a-pull-request
 
 POST /repos/:owner/:repo/pulls
 
 {
   "title": "Amazing new feature",
   "body": "Please pull this in!",
   "head": "octocat:new-feature",
   "base": "master"
 }

Maybe try [Go Github](https://github.com/google/go-github)
