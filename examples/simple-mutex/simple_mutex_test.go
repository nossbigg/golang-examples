package examples

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMutex(t *testing.T) {
	values := make(chan int, 10)
	exit := make(chan struct{})

	go SimpleMutex(values, exit)

	time.Sleep(10 * time.Millisecond)
	close(exit)

	receivedValues := []int{}
	for i := 0; i < 10; i++ {
		receivedValues = append(receivedValues, <-values)
	}

	expectedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.ElementsMatch(t, expectedValues, receivedValues)
}
