package cmd

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/sambasivareddy-ch/notebook_cli/data"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete a notes based on title",
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

func init() {
	// Appending Delete Command as child to Root
	rootCmd.AddCommand(deleteCmd)
}

func delete() {
	// Get all the titles to pass them to SELECT prompt
	titles := data.GetAllTitlesFromNoteBooksTable()

	// Creating SELECT prompt
	selectTitlePrompt := &promptui.Select{
		Label: "Select Title you wanted to delete",
		Items: titles,
	}

	// get the selected title
	_, selectedTitle, err := selectTitlePrompt.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}

	data.DeleteNotesWithGivenTitle(selectedTitle)
}
