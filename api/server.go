package api

import (
	"database/sql"
	"fmt"
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
	port = ":" + port
	fmt.Println("Serving On:", port)
	// for HEROKU: ":"+os.Getenv("PORT")
	err := http.ListenAndServe(port, router)
	if err != nil { // should I log fatal??
		fmt.Println(err)
	}
}
