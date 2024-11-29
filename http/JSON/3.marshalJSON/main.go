package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {

	bytes := make([][]byte, 0)

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, data)
	}

	return bytes, nil
}

/*
type Board struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TeamId   int    `json:"team"`
	TeamName string `json:"team_name"`
}

board := Board{
	Id:       1,
	Name:     "API",
	TeamId:   9001,
	TeamName: "Backend",
}

data, err := json.Marshal(board)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(data)) */
// {"id":1,"name":"API","team":9001,"team_name":"Backend"}
