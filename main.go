package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "mypassword"

	result, err := executeCypherQueries(context.Background(), uri, username, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}

func readQueriesFromFile(filepath string) ([]string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Split the file content into individual queries based on a separator
	queries := strings.Split(string(content), ";")

	// Remove any empty queries
	var cleanedQueries []string
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {
			cleanedQueries = append(cleanedQueries, query)
		}
	}

	return cleanedQueries, nil
}

func executeCypherQueries(ctx context.Context, uri, username, password string) (string, error) {
	queries, err := readQueriesFromFile("queries.cypher")
	if err != nil {
		return "", err
	}

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close()

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session := driver.NewSession(sessionConfig)
	defer session.Close()

	for _, query := range queries {
		parameters := map[string]interface{}{}
		result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
			return transaction.Run(query, parameters)
		})
		if err != nil {
			fmt.Println("Error executing query:", err, result)
			return "", err
		}
	}

	return "Queries executed successfully", nil
}
