package examples

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAtomicAdd(t *testing.T) {
	values := make(chan int32, 10)
	exit := make(chan struct{})

	go SimpleAtomicAdd(values, exit)

	time.Sleep(10 * time.Millisecond)
	close(exit)

	receivedValues := []int32{}
	for i := 0; i < 10; i++ {
		receivedValues = append(receivedValues, <-values)
	}

	expectedValues := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.ElementsMatch(t, expectedValues, receivedValues)
}
