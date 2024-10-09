package db

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var KNOWLEDGE_BASE_IN_USE string

type SearchResult struct {
	ElementID string
	Text      string
	Filename  string
	Score     float64
}

func Search(input string, embedding []float64, tableName string) []SearchResult {
	var embeddingStringArray []string
	for _, num := range embedding {
		embeddingStringArray = append(embeddingStringArray, strconv.FormatFloat(num, 'f', -1, 64))
	}
	embeddingString := fmt.Sprintf("[%s]", strings.Join(embeddingStringArray, ", "))

	query := fmt.Sprintf(`SELECT element_id, text, filename,
	list_cosine_similarity(embeddings, %v) as score
	FROM %s.%s.%s
	order by score desc limit 10;`, embeddingString, localDatabase, localSchema, tableName)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var searchResults []SearchResult

	for rows.Next() {
		var result SearchResult
		err = rows.Scan(&result.ElementID, &result.Text, &result.Filename, &result.Score)
		if err != nil {
			log.Fatal(err)
		}
		searchResults = append(searchResults, result)

	}

	return searchResults
}
