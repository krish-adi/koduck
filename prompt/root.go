package prompt

import (
	"net/http"
)

// ollamaClient variable for the reusable HTTP client
var ollamaClient *http.Client

// ollamaBaseURL variable for the reusable base URL
var ollamaBaseURL string

var MODEL_IN_USE_COMPLETION string

// Initialize the HTTP client for the package
func InitClients() {
	MODEL_IN_USE_COMPLETION = "llama3.2"
	ollamaClient = &http.Client{}
	ollamaBaseURL = "http://localhost:11434/api"
}
