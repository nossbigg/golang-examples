package main

import (
	"fmt"
	"time"

	examples "nossbigg.com/golangexamples/examples/simple-threadpool"
)

func main() {
	exit := make(chan struct{})

	processedElements := make([]int, 0)
	processedElementsChannel := make(chan int)

	go func() {
		for {
			select {
			case v := <-processedElementsChannel:
				processedElements = append(processedElements, v)
			case <-exit:
				return
			}
		}
	}()

	options := &examples.SimpleThreadpoolOptions{
		NThreads:                 2,
		GeneratorSleepMs:         100,
		LinebreakSleepMs:         1000,
		WorkerSleepMs:            200,
		EnableLinebreaks:         true,
		Printf:                   fmt.Printf,
		Exit:                     exit,
		Debug:                    true,
		ProcessedElementsChannel: processedElementsChannel,
	}
	go examples.SimpleThreadpool(*options)

	time.Sleep(time.Duration(2) * time.Second)
	close(exit)

	fmt.Println("closed:")
	time.Sleep(time.Duration(3) * time.Second)

	fmt.Printf("processedElements: ")
	fmt.Print(processedElements)
	fmt.Println()
	fmt.Println("main: exited")
}
