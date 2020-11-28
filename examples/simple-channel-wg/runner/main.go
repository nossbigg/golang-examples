package main

import (
	"fmt"

	examples "nossbigg.com/golangexamples/examples/simple-channel-wg"
)

func main() {
	returnChannel := make(chan []int)

	go examples.SimpleChannelWg(10, returnChannel)

	result := <-returnChannel
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
