package main
/*
Characteristics of a Golang test function:

The first and only parameter needs to be t *testing.T
It begins with the word Test followed by a word or phrase starting with a capital letter.
(usually the method under test i.e. TestValidateClient)
Calls t.Error or t.Fail to indicate a failure (I called t.Errorf to provide more details)
t.Log can be used to provide non-failing debug information
Must be saved in a file named something_test.go such as: addition_test.go
 */
import "testing"


func TestBuildRequest(t *testing.T) {
	url := "https://hahanotreal.com"
	jiraUserId := "notme"
	jiraPassword := "nope"

	_, req := BuildRequest(url, jiraUserId, jiraPassword)

	if req.URL.String() != url {
		t.Fatalf("BuildRequest(%s, %s, %s) = %v, want %v", url, jiraUserId, jiraPassword, req.URL.String(), url)
	} else {
		t.Log("BuildRequest URL correct" )
	}

	if req.Header.Get("Accept") != "application/json" {
		t.Fatalf("BuildRequest(%s, %s, %s) = %v, want %v", url, jiraUserId, jiraPassword, req.Header.Get("Accept"), "application/json")
	} else {
		t.Log("BuildRequest Accept json header correct" )
	}

	if req.Header.Get("Authorization") != "Basic bm90bWU6bm9wZQ==" {
		t.Fatalf("BuildRequest(%s, %s, %s) = %v, want %v", url, jiraUserId, jiraPassword, req.Header.Get("Authorization"), "application/json")
	} else {
		t.Log("BuildRequest Authorization header correct" )
	}
}
func TestGetJiraResponse(t *testing.T) {
	//jiraResponse := GetJiraResponse(jiraClient, req)

}