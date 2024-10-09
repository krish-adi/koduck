package prompt

import (
	"net/http"
)

// ollamaClient variable for the reusable HTTP client
var ollamaClient *http.Client

// ollamaBaseURL variable for the reusable base URL
var ollamaBaseURL string

// Initialize the HTTP client for the package
func InitClients() {
	ollamaClient = &http.Client{}
	ollamaBaseURL = "http://localhost:11434/api"
}
