package main

import (
	"fmt"

	"github.com/polypmer/mappyr/api"
	"github.com/polypmer/mappyr/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error DB", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		fmt.Println("Creation error", err)
	}

	api.Serve(db)
}
