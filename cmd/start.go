package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/krish-adi/koduck/db"
	"github.com/krish-adi/koduck/prompt"
)

func init() {
	// Add the session command to the root command
	rootCmd.AddCommand(startSessionCmd)
}

// Command that starts a session
var startSessionCmd = &cobra.Command{
	Use:   "start",
	Short: "Start koduck",
	Run: func(cmd *cobra.Command, args []string) {
		startSession() // Start the interactive session
	},
}

// Function that starts the session
func startSession() {
	reader := bufio.NewReader(os.Stdin) // Reading input from the user

	for {
		fmt.Print("\033[38;5;208mkoduck> \033[0m")

		// Read user input until EOF (Ctrl+D)
		input, err := reader.ReadString('\n')

		// Check for EOF (Ctrl+D)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\n\nExiting koduck...")
				break
			}
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim newline and extra spaces from input
		input = strings.TrimSpace(input)

		// Check if the input is ".list"
		if input == ".list" {
			db.List()
			continue
		} else if strings.HasPrefix(input, ".pull ") {
			parts := strings.Fields(input)
			if len(parts) == 2 {
				tableName := parts[1]
				db.Pull(tableName)
				continue
			} else {
				fmt.Println("Invalid input. Please provide a knowledge base name.")
				continue
			}
		} else if strings.HasPrefix(input, ".drop ") {
			parts := strings.Fields(input)
			if len(parts) == 2 {
				tableName := parts[1]
				db.Drop(tableName)
				continue
			} else {
				fmt.Println("Invalid input. Please provide a knowledge base name.")
				continue
			}
		} else {
			_, err = prompt.LLM(input, "llama3")
			if err != nil {
				fmt.Print("\n")
				fmt.Println("Error generating answer:", err)
				fmt.Print("\n")
				fmt.Print("\n")
				continue
			}
		}

		fmt.Print("\n")
		fmt.Print("\n")
	}
}
