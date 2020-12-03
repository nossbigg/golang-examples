package main

import (
	"fmt"
	"time"

	examples "nossbigg.com/golangexamples/examples/simple-mutex"
)

func main() {
	nValues := 20
	values := make(chan int, nValues)
	exit := make(chan struct{})

	go examples.SimpleMutex(values, exit)

	time.Sleep(100 * time.Millisecond)
	close(exit)

	for i := 0; i < nValues; i++ {
		fmt.Printf("%d ", <-values)
	}
}
