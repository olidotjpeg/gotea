package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Server is live on port 8000")

	var err error

	Database, err = sql.Open("sqlite3", "./gotea.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer Database.Close()

	HandleRequests()
}
