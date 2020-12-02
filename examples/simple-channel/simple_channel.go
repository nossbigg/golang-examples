package examples

// SimpleChannel ...
func SimpleChannel(times int) []int {
	outputChannel := make(chan int)

	for i := 0; i < times; i++ {
		go timesTwo(i+1, outputChannel)
	}

	var result []int
	for i := 0; i < times; i++ {
		value := <-outputChannel
		result = append(result, value)
	}

	return result
}

func timesTwo(value int, outputChannel chan<- int) {
	returnValue := value * 2
	outputChannel <- returnValue
}
