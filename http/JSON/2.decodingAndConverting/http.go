package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// res is a successful `http.Response`

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&issues); err != nil {
		fmt.Println("error decoding response body")
		return nil, err
	}

	return issues, nil
}
