package main

import (
	"fmt"

	"github.com/sambasivareddy-ch/notebook_cli/cmd"
	"github.com/sambasivareddy-ch/notebook_cli/data"
)

func main() {
	// Creating Database on root command
	if err := data.CreateDatabase(); err != nil {
		fmt.Println("Unable to open database")
	}
	cmd.Execute()
}
