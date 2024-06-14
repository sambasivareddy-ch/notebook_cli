package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Root Command",
	Short: "Entry point to Notebook CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Notebook")
	},
}

func Execute() {
	rootCmd.Execute()
}
