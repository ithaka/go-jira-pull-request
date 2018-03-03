package main

import (
	"context"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"os"
	"fmt"
)

/*
 https://developer.github.com/v3/pulls/#create-a-pull-request
type NewPullRequest struct {
	Title               *string `json:"title,omitempty"`
	Head                *string `json:"head,omitempty"`
	Base                *string `json:"base,omitempty"`
	Body                *string `json:"body,omitempty"`
	Issue               *int    `json:"issue,omitempty"`
	MaintainerCanModify *bool   `json:"maintainer_can_modify,omitempty"`
}
 */


func main() {
	ctx := context.Background()
	token := getEnvOrDie("GO_JIRA_PULL_REQUEST_AUTH_TOKEN")
	//GO_JIRA_PULL_REQUEST_AUTH_TOKEN
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	//repos, _, _ := client.Repositories.List(ctx, "StevenACoffman", nil)
	//fmt.Println(repos)


	input := &github.NewPullRequest{Title: github.String("titleiscool"), Head: github.String("testing"), Body: github.String("So this is a description")}
	pull, response, err := client.PullRequests.Create(context.Background(), "ithaka", "fluent-bit-json-test", input)
	if err != nil {
		fmt.Errorf("PullRequests.Create returned error: %v", err)
	}
	fmt.Print(pull)
	fmt.Print(response.String())

}

func getEnvOrDie(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic("No github personal access token in env "+key)
}