package main

import (
	"fmt"

	db "nossbigg.com/golangexamples/examples/db-common"
	examples "nossbigg.com/golangexamples/examples/db-select-items"
)

func main() {
	conn := db.GetDbConnection()
	items := examples.DbSelectItems(conn)

	fmt.Println("Items: ")
	for _, item := range items {
		fmt.Println(item)
	}
}
