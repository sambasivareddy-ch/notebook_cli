**NoteBook CLI in Go**

Repository: https://github.com/sambasivareddy-ch/notebook_cli

Build: To build the app use 
<code>go build .</code>
This will create a build file/executable file notebook_cli

***Command***
- InitTable   to initialize the notebook table
- completion  Generate the autocompletion script for the specified shell
- create      To create new notes
- delete      To delete a notes based on title
- help        Help about any command
- modify      Modifies the notes based on title
- view        Displays all the notes existed

***Usage***
./notebook_cli -> Will creates a database called sqlite-database.db
./notebook_cli InitTable -> Will creates a NoteBooks table in sqlite-database.db file

Now the CLI app is ready, now you can play around with various above command
./notebook_cli view -> Gives all the notes in the table
./notebook_cli create -> gives the prompt to enter the title & notes. Once entered a new notes will be created
./notebook_cli delete -> gives the prompt to select the title and deletes a notes from the table based on that select title
./notebook_cli modify -> gives the prompt to select the title and a prompt to enter new notes & update the table
./notebook_cli --help -> List out all the available command