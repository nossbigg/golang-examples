package main

import (
	db "nossbigg.com/golangexamples/examples/db-common"
	examples "nossbigg.com/golangexamples/examples/db-delete-item"
)

func main() {
	conn := db.GetDbConnection()
	examples.DbDeleteItem(conn, 4)
}
