package examples

import (
	"fmt"

	"github.com/jackc/pgx"
)

// DbDeleteItem ...
func DbDeleteItem(conn *pgx.Conn, id int) {
	defer conn.Close()

	_, err := conn.Query("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
