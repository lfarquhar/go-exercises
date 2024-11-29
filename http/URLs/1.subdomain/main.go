package main

import (
	"log"
)

func main() {
	issues, err := getIssues(subdomain, domain)
	if err != nil {
		log.Fatalf("error getting issues data: %v", err)
	}
	logIssues(issues)
}
