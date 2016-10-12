package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

// TODO:
// Define a global database variablei
var db *sql.DB

// Serve connection

func Serve(connection *sql.DB, port string) { // pass in connection
	// set connection to global
	db = connection
	fmt.Println("Serving API on port 8080")

	router := NewRouter()
	fmt.Println("Should be", port)
	err := http.ListenAndServe(os.Getenv("PORT"), router)
	if err != nil { // should I log fatal??
		fmt.Println(err)
	}
}
