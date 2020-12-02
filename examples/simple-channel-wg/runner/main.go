package main

import (
	"fmt"

	examples "nossbigg.com/golangexamples/examples/simple-channel-wg"
)

func main() {
	result := examples.SimpleChannelWg(10)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
