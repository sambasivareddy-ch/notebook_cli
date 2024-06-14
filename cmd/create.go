package cmd

import (
	"errors"
	"log"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/sambasivareddy-ch/notebook_cli/data"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "To create new notes",
	Run: func(cmd *cobra.Command, args []string) {
		getDetailsFromUser()
	},
}

func init() {
	// Appending Create Command as child to Root
	rootCmd.AddCommand(createCmd)
}

func getDetailsFromUser() {
	var err error
	var title, notes string

	// validate user input
	validator := func(userInput string) error {
		if len(userInput) <= 0 {
			return errors.New("length should be greater than 0")
		}
		return nil
	}

	// title & notes prompt to get data from users
	titlePrompt := promptui.Prompt{
		Label:    "Title",
		Validate: validator,
	}

	notesPrompt := promptui.Prompt{
		Label:    "Notes",
		Validate: validator,
	}

	// get the title entered by user
	title, err = titlePrompt.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	// get the notes entered by user
	notes, err = notesPrompt.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	data.InsertIntoNotebookTable(title, notes, time.Now())
}
