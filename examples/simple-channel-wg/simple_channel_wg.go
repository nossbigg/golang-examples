package examples

import (
	"sync"
)

// SimpleChannelWg ...
func SimpleChannelWg(times int) []int {
	wg := &sync.WaitGroup{}
	wg.Add(times)

	outputChannel := make(chan int, times)

	for i := 0; i < times; i++ {
		go timesTwo(i+1, outputChannel, wg)
	}

	wg.Wait()

	var result []int
	for i := 0; i < times; i++ {
		value := <-outputChannel
		result = append(result, value)
	}

	return result
}

func timesTwo(value int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	returnValue := value * 2
	result <- returnValue
}
