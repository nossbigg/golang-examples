package examples

import (
	"fmt"
	"os"

	pgx "github.com/jackc/pgx"
)

// ItemRow ...
type ItemRow struct {
	Id   int
	Name string
}

// GetDbConnection ...
func GetDbConnection() *pgx.Conn {
	connectionConfig := &pgx.ConnConfig{
		Database: "golangexamples",
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "example",
	}

	// Note: Don't define your database url in code in prod!
	conn, err := pgx.Connect(*connectionConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
