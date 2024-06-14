package cmd

import (
	"github.com/sambasivareddy-ch/notebook_cli/data"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "InitTable",
	Short: "to initialize the notebook table",
	Run: func(cmd *cobra.Command, args []string) {
		data.CreateTable()
	},
}

func init() {
	// Appending Init Command as child to Root
	rootCmd.AddCommand(initCmd)
}
