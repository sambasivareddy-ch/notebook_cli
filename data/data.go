package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/*
Global Database instance which will be used to perform all the operations
*/
var database *sql.DB

/*
CreateDatabase() - Create a new database file called sqlite-database.db within the same project
- We are actually mimicing the Database, we don't have server now.
*/
func CreateDatabase() error {
	var err error
	// Opens a Database given by driver & returns it
	database, err = sql.Open("sqlite3", "./sqlite-database.db")

	// If error occurs while opening database
	if err != nil {
		log.Fatalf("Unable to open the database: %v", err)
	}

	return database.Ping()
}

/*
CreateTable() - Creates a NoteBooks(title primary key, notes, created_on) table in give sqlite-database.db file
*/
func CreateTable() {
	notebookTableCommand := `CREATE TABLE IF NOT EXISTS NOTEBOOKS (
		"TITLE" VARCHAR(30) PRIMARY KEY,
		"NOTES" TEXT,
		"CREATED_ON" TIMESTAMP
	)`

	// Preparing executable prepared statement
	stmt, err := database.Prepare(notebookTableCommand)

	if err != nil {
		log.Fatal(err.Error())
	}

	// Executing the prepared statement & handling error if occurs
	if _, err1 := stmt.Exec(); err1 != nil {
		log.Fatal(err1.Error())
	}

	fmt.Println("Notebook Table Created Successfully :)")
}

/*
InsertIntoNotebookTable() - Inserts the passed title, notes, created_on in the table NoteBooks
*/
func InsertIntoNotebookTable(title string, notes string, created_on time.Time) {
	// Here, ? represents placeholders in the prepared statement
	insertCommand := `INSERT INTO NOTEBOOKS VALUES(?,?,?)`

	stmt, err := database.Prepare(insertCommand)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Passing the values to the prepared statement generated above at the time of execution
	_, err1 := stmt.Exec(title, notes, created_on)

	if err1 != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Notes with title: %s is created successfully\n", title)
}

/*
UpdateNotesWithGivenTitle() - Updates a row/notes with given newNotes text based on selected title
*/
func UpdateNotesWithGivenTitle(title string, newNotes string) {
	// Here, ? represents placeholders in the prepared statement
	updateCommand := `UPDATE NOTEBOOKS SET NOTES = ?, CREATED_ON = ? WHERE TITLE = ?`

	stmt, err := database.Prepare(updateCommand)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Passing the values to the prepared statement generated above at the time of execution
	_, err1 := stmt.Exec(newNotes, time.Now(), title)

	if err1 != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Notes with title: %s is updated successfully\n", title)
}

/*
DeleteNotesWithGivenTitle() - Deletes a row/notes with the selected title
*/
func DeleteNotesWithGivenTitle(title string) {
	// Here, ? represents placeholders in the prepared statement
	deleteCommand := `DELETE FROM NOTEBOOKS WHERE TITLE = ?`

	stmt, err := database.Prepare(deleteCommand)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Passing the values to the prepared statement generated above at the time of execution
	_, err1 := stmt.Exec(title)

	if err1 != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Notes with title: %s is deleted successfully\n", title)
}

/*
ShowNotes() - Show all the tuples/rows in the notebooks table
*/
func ShowNotes() {
	selectCommand := "SELECT * FROM NOTEBOOKS"

	stmt, err := database.Prepare(selectCommand)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Query() returns all the tuples returned after executing the
	// prepared statement & throws error if any problem occurs
	rows, err1 := stmt.Query()
	if err1 != nil {
		log.Fatalf(err1.Error())
	}

	// Iterating throw each row returned by Query() and printing the row in the formatted way
	for rows.Next() {
		var title, notes string
		var created_at time.Time

		rows.Scan(&title, &notes, &created_at)

		rowString := fmt.Sprintf("Title: %s | Notes: %s | Created At: %v\n", title, notes, created_at)

		fmt.Println(rowString)
	}
}

/*
GetAllTitlesFromNoteBooksTable() - Returns all the titles in the given NoteBooks table
*/
func GetAllTitlesFromNoteBooksTable() []string {
	selectQuery := "SELECT title FROM NOTEBOOKS"

	stmt, err := database.Prepare(selectQuery)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Query() returns all the tuples returned after executing the
	// prepared statement & throws error if any problem occurs
	rows, err1 := stmt.Query()
	if err1 != nil {
		log.Fatalf(err1.Error())
	}

	// Create a slice of strings & store all the titles in it & return them
	titles := []string{}
	for rows.Next() {
		var title string
		rows.Scan(&title)

		titles = append(titles, title)
	}

	return titles
}
