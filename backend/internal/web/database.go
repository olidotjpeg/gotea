package web

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

const connString = "postgresql://docker:postgres123@postgres:5432/gotea?sslmode=disable"

func databaseInit() {
	var err error
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v\n", err)
	}

	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
}
