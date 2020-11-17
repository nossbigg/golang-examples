package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLettersCount(t *testing.T) {
	result := lettersCount("abb")

	expectedResult := map[string]int{
		"a": 1,
		"b": 2,
	}
	assert.Equal(t, result, expectedResult)
}
