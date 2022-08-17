package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/olidotjpeg/gotea/packages/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Server is live on port 8000")

	var err error

	handlers.Database, err = sql.Open("sqlite3", "./gotea.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer handlers.Database.Close()

	handlers.HandleRequests()
}
