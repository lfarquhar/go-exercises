package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getIssues(url string) ([]Issue, error) {
	data, err := getIssueData(url)

	var issues []Issue
	if err = json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}
