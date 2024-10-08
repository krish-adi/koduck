package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Add a variable to store the version
var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "koduck",
	Short: "koduck, a tool for managing knowledge bases.",
	Long:  "koduck, a tool for managing knowledge bases.",
	// Add a Run function to handle the version flag
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("koduck version is %s\n", version)
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	// Remove the default 'completion' command if it exists
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add the version flag
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
}
