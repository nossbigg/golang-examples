package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleChannel(t *testing.T) {
	result := SimpleChannel(5)

	expectedValues := []int{2, 4, 6, 8, 10}
	for _, v := range expectedValues {
		assert.Contains(t, result, v)
	}
}
