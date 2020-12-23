package main

import (
	db "nossbigg.com/golangexamples/examples/db-common"
	examples "nossbigg.com/golangexamples/examples/db-insert-item"
)

func main() {
	conn := db.GetDbConnection()
	examples.DbInsertItem(conn, "item-name")
}
