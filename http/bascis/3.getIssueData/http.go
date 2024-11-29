package main

import (
	"fmt"
	"io"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close() // defer res.Body.Close() ensures that the response body is properly closed after reading. Not doing so can cause memory issues.

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}
