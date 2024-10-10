package prompt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func Completion(message string, context []string) (CompletionResponse, error) {
	// API endpoint
	url := ollamaBaseURL + "/generate"

	prompt := `<|begin_of_text|><|start_header_id|>system<|end_header_id|>
You are a helpful assistant. Given the following context: ` +
		strings.Join(context, "\n") + ` answer the following question from 
		the user. Answer in brief, and use the keywrods provided in the context. <|eot_id|><|start_header_id|>user<|end_header_id|> ` +
		message + ` <|eot_id|><|start_header_id|>assistant<|end_header_id|>`

	// Create the request payload
	completionReq := CompletionRequest{
		Model:  MODEL_IN_USE_COMPLETION,
		Prompt: prompt,
	}

	// Marshal the request to JSON
	jsonData, err := json.Marshal(completionReq)
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
