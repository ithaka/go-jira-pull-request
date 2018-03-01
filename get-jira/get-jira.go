package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"os"
	"encoding/json"
)

type JIRAResponse struct {
	Key    string `json:"key"`
	Fields struct {
		Description          string      `json:"description"`
		Summary          string        `json:"summary"`
	} `json:"fields"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	jiraUserId := getEnv("JIRA_LOGIN", "login")
	jiraPassword := getEnv("JIRA_PASSWORD", "password")

	url := "https://jira.jstor.org/rest/api/2/issue/CORE-5339"

	jiraClient := http.Client{
		Timeout: time.Second * 15, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.SetBasicAuth(jiraUserId, jiraPassword)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type","application/json")
	res, getErr := jiraClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jiraResponse := JIRAResponse{}
	jsonErr := json.Unmarshal(body, &jiraResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(jiraResponse.Key) 
	fmt.Println(jiraResponse.Fields.Summary)
	fmt.Println(jiraResponse.Fields.Description)
}