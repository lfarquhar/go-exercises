package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Title string
}

func getIssues(subdomain string, domain string) ([]Issue, error) {

	fullDomain := domain

	if len(subdomain) > 0 {
		fullDomain = subdomain + "." + domain
	}

	res, err := http.Get("https://" + fullDomain + "/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func logIssues(issues []Issue) {
	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}
