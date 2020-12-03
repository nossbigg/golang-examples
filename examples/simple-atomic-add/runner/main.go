package main

import (
	"fmt"
	"time"

	examples "nossbigg.com/golangexamples/examples/simple-atomic-add"
)

func main() {
	nValues := 20
	values := make(chan int32, nValues)
	exit := make(chan struct{})

	go examples.SimpleAtomicAdd(values, exit)

	time.Sleep(100 * time.Millisecond)
	close(exit)

	for i := 0; i < nValues; i++ {
		fmt.Printf("%d ", <-values)
	}
}
