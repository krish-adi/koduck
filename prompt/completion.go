package prompt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CompletionRequest represents the request structure for the generation API
type CompletionRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// CompletionResponse represents the response structure from the generation API
type CompletionResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
	Context   []int     `json:"context"`
}

func LLM(message string, model string) (CompletionResponse, error) {
	// API endpoint
	url := ollamaBaseURL + "/generate"

	// Create the request payload
	llmReq := CompletionRequest{
		Model:  model,
		Prompt: message,
	}

	// Marshal the request to JSON
	jsonData, err := json.Marshal(llmReq)
	if err != nil {
		return CompletionResponse{}, fmt.Errorf("error marshaling request: %w", err)
	}

	// Create a new HTTP POST request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return CompletionResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// Set the appropriate headers for a JSON payload
	req.Header.Set("Content-Type", "application/json")

	resp, err := ollamaClient.Do(req)
	if err != nil {
		return CompletionResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Create a decoder to read the JSON stream
	decoder := json.NewDecoder(resp.Body)

	var completionResponse CompletionResponse

	for {

		if err := decoder.Decode(&completionResponse); err != nil {
			if err == io.EOF {
				break
			}
			return CompletionResponse{}, fmt.Errorf("error decoding response: %w", err)
		}

		fmt.Print(completionResponse.Response)

		if completionResponse.Done {
			break
		}
	}

	return completionResponse, nil
}
