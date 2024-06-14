package cmd

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/sambasivareddy-ch/notebook_cli/data"
	"github.com/spf13/cobra"
)

var modifyCommand = &cobra.Command{
	Use:   "modify",
	Short: "Modifies the notes based on title",
	Run: func(cmd *cobra.Command, args []string) {
		modify()
	},
}

func init() {
	// Appending Modify Command as child to Root
	rootCmd.AddCommand(modifyCommand)
}

func modify() {
	// validate function to validate user entered data
	validator := func(userInput string) error {
		if len(userInput) <= 0 {
			return errors.New("length should be greater than 0")
		}
		return nil
	}

	// Get all the titles to pass them to SELECT prompt
	titles := data.GetAllTitlesFromNoteBooksTable()

	// Creating SELECT prompt
	selectTitlePrompt := &promptui.Select{
		Label: "Select Title you wanted to update",
		Items: titles,
	}

	// gets the selected title from the prompted titles to the users
	_, selectedTitle, err := selectTitlePrompt.Run()

	if err != nil {
		log.Fatalf(err.Error())
	}

	// New prompt to ask user for modified notes
	newNotesPrompt := &promptui.Prompt{
		Label:    "Enter new notes",
		Validate: validator,
	}

	// Run the prompt & get user entered notes
	newNotes, err1 := newNotesPrompt.Run()

	if err1 != nil {
		log.Fatalf(err.Error())
	}

	data.UpdateNotesWithGivenTitle(selectedTitle, newNotes)
}
