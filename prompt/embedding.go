package prompt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// EmbedRequest represents the request structure for the embedding API
type EmbedRequest struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
}

// EmbedResponse represents the response structure from the embedding API
type EmbedResponse struct {
	Model           string      `json:"model"`
	Embeddings      [][]float64 `json:"embeddings"`
	TotalDuration   int64       `json:"total_duration"`
	LoadDuration    int64       `json:"load_duration"`
	PromptEvalCount int         `json:"prompt_eval_count"`
}

func Embedding(input []string, model string) (EmbedResponse, error) {
	// API endpoint
	url := ollamaBaseURL + "/embed"

	// Create the request payload
	embedReq := EmbedRequest{
		Model: model,
		Input: input,
	}

	// Marshal the request to JSON
	jsonData, err := json.Marshal(embedReq)
	if err != nil {
		return EmbedResponse{}, fmt.Errorf("error marshaling request: %w", err)
	}

	// Create a new HTTP POST request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return EmbedResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// Set the appropriate headers for a JSON payload
	req.Header.Set("Content-Type", "application/json")

	resp, err := ollamaClient.Do(req)
	if err != nil {
		return EmbedResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var embedResponse EmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embedResponse); err != nil {
		return EmbedResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return embedResponse, nil
}
