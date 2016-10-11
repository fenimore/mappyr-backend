package api

import "fmt"

// TODO:
// Define a global database variablei
// var db

// Serve connection

func Serve() { // pass in connection
	// set connection to global

	fmt.Println("Serving API on port 8080")

	router := NewRouter()

	err := http.ListenAndServer(":8080", router)
	if err != nil { // should I log fatal??
		fmt.Println(err)
	}
}
