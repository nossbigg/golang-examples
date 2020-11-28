package main

import (
	"fmt"

	examples "nossbigg.com/golangexamples/examples/simple-channel"
)

func main() {
	result := examples.SimpleChannel(10)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
