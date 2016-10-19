package api

import (
	"database/sql"
	"fmt"
	"github.com/polypmer/mappyr-backend/database"
	"net/http"
)

// TODO:
// Define a global database variablei
var db *sql.DB

// Serve connection

func Serve(connection *sql.DB, port string) { // pass in connection
	// set connection to global
	db = connection
	// get new router
	router := NewRouter()
	// get port
	//port = ":" + port
	fmt.Println("Serving On:", port)

	e := database.MockVote(db) // database.MockUsers(db)
	if e != nil {
		fmt.Println(e)
	}
	e = database.MockUsers(db)
	if e != nil {
		fmt.Println(e)
	}
	e = database.TallyVotes(db)
	if e != nil {
		fmt.Println(e)
	}

	// for HEROKU: ":"+os.Getenv("PORT")
	err := http.ListenAndServe(port, router)
	if err != nil { // should I log fatal??
		fmt.Println(err)
	}
}
