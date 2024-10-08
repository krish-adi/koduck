package prompt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type StreamResponse struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Response  string    `json:"response"`
	Done      bool      `json:"done"`
	Context   []int     `json:"context"`
}

func LLM(message string, model string) (StreamResponse, error) {
	// API endpoint
	url := "http://localhost:11434/api/generate"

	// JSON payload
	jsonData := []byte(fmt.Sprintf(`{
		"model": "%s",
		"prompt": "%s"
	}`, model, message))

	// Create a new HTTP POST request with the payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return StreamResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	// Set the appropriate headers for a JSON payload
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return StreamResponse{}, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Create a decoder to read the JSON stream
	decoder := json.NewDecoder(resp.Body)

	var streamResponse StreamResponse

	for {

		if err := decoder.Decode(&streamResponse); err != nil {
			if err == io.EOF {
				break
			}
			return StreamResponse{}, fmt.Errorf("error decoding response: %w", err)
		}

		fmt.Print(streamResponse.Response)

		if streamResponse.Done {
			break
		}
	}

	return streamResponse, nil
}
