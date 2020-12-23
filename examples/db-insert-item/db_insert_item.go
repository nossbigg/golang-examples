package examples

import (
	"fmt"

	"github.com/jackc/pgx"
)

// DbInsertItem ...
func DbInsertItem(conn *pgx.Conn, name string) {
	defer conn.Close()

	_, err := conn.Query("INSERT INTO items(name) VALUES ($1)", name)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
