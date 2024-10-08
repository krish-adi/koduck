package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List knowledge bases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1. koduck")
		fmt.Println("2. koduck2")
	},
}
