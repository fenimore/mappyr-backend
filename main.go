package main

import (
	"fmt"

	"github.com/polypmer/sunken/api"
	"github.com/polypmer/sunken/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error DB", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		fmt.Println(err)
	}
	api.Serve(db)
}
