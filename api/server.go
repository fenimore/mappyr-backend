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

func Serve(connection *sql.DB) { // pass in connection
	// set connection to global
	db = connection
	fmt.Println("Serving API on port 8080")

	router := NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil { // should I log fatal??
		fmt.Println(err)
	}
}
