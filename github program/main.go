package main

import (
	"context" // Have imported the packages  --> handles the contexts
	"fmt"     // used for formatting the output
	"log"     // used for logging the errors and messages
	"os"      // used for interacting with the operating system

	"github.com/google/go-github/v69/github" // go-github library
	"golang.org/x/oauth2"                    // used to handle the token based authentication
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN") // Environmental variables
	if githubToken == "" {
		log.Fatal("Error: GITHUB_TOKEN environment variable not set") // checking if the token is empty if the error shows up it will stop execution
	}

	ctx := context.Background()     // creating a background context
	ts := oauth2.StaticTokenSource( //  used this OAuth2 token source using the GitHub token for authentication
		&oauth2.Token{AccessToken: githubToken}, // access tokens are obtained from environmental variable
	)
	tc := oauth2.NewClient(ctx, ts) // creating the HTTP client from token source
	client := github.NewClient(tc)  // created a new github client using oauth2
	org := "MachaniRobotics"

	opt := &github.RepositoryListByOrgOptions{Type: "all"}        // fetching all repositories
	repos, _, err := client.Repositories.ListByOrg(ctx, org, opt) // handles all the errors occur during API call
	if err != nil {
		log.Fatalf("Error fetching repositories: %v", err) // if the error occurs the execution will stop
	}

	fmt.Println("Repositories of", org) // printing the organisation name
	for i, repo := range repos {        //
		fmt.Println(i+1, *repo.Name, ":-", *repo.HTMLURL) // printing the respositories with number, name and HTTP url
	}
}
