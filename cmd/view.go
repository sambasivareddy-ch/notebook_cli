package cmd

import (
	"github.com/sambasivareddy-ch/notebook_cli/data"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Displays all the notes existed",
	Run: func(cmd *cobra.Command, args []string) {
		data.ShowNotes()
	},
}

func init() {
	// Appending View Command as child to Root
	rootCmd.AddCommand(viewCmd)
}
