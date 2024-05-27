package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
)

var (
	UNAMEDB string = "postgres"
	PASSDB  string = "postgres123"
	HOSTDB  string = "postgres"
	DBNAME  string = "teas"
)

func main() {
	fmt.Println("Server is live on port 8000")

	// var err error

	connString := "postgresql://postgres:postgress123@localhost/teas"
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return
	}
	defer conn.Close(context.Background())

	// Perform a simple query
	var result string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL!'").Scan(&result)
	if err != nil {
		fmt.Println("Query failed:", err)
		return
	}

	fmt.Println(result)

	// web.Database, err = sql.Open("sqlite3", "./gotea.sqlite")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer web.Database.Close()

	// web.HandleRequests()
}
