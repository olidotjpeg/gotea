package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olidotjpeg/backend/internal/web"
)

var (
	UNAMEDB string = "postgres"
	PASSDB  string = "postgres123"
	HOSTDB  string = "postgres"
	DBNAME  string = "teas"
)

func main() {
	fmt.Println("Server is live on port 8000")

	web.HandleRequests()
}
