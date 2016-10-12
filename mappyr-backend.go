package mappyr-backend

import (
	"fmt"

	"github.com/polypmer/mappyr-backend/api"
	"github.com/polypmer/mappyr-backend/database"
)

func main() {
	db, err := database.InitDB()
	defer db.Close() // This ought to have been a while ago
	if err != nil {
		fmt.Println("Error DB", err)
	}

	err = database.CreateTable(db)
	if err != nil {
		fmt.Println("Creation error", err)
	}

	api.Serve(db)
}
