package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleChannelWg(t *testing.T) {
	resultChannel := make(chan []int)

	go SimpleChannelWg(10, resultChannel)

	result := <-resultChannel

	expectedValues := []int{2, 4, 6, 8, 10}
	for _, v := range expectedValues {
		assert.Contains(t, result, v)
	}
}
