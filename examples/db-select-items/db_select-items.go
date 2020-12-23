package examples

import (
	"github.com/jackc/pgx"

	db "nossbigg.com/golangexamples/examples/db-common"
)

// DbSelectItems ...
func DbSelectItems(conn *pgx.Conn) []db.ItemRow {
	defer conn.Close()

	items := []db.ItemRow{}

	rows, err := conn.Query("SELECT id, name FROM items;")
	if err != nil {
		return items
	}

	for rows.Next() {
		var id int
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			return items
		}

		newItem := &db.ItemRow{
			Id:   id,
			Name: name,
		}
		items = append(items, *newItem)
	}

	return items
}
