package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/shamimgeek/repos")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{}

	jsonErr := json.Unmarshal(body, &results)
	if err != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("%-50s\t%-10s\t%10s\t%10s\n", "Project Name", "Language", "Star Count", "Fork Count")
	for _, result := range results {
		fmt.Printf("%-50s\t%-10v\t%10.f\t%10.f\n", result["name"], result["language"], result["stargazers_count"], result["forks_count"])
	}

}
