package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLettersCount(t *testing.T) {
	result := lettersCount("abb")

	expectedResult := lettersCountMap{
		"a": 1,
		"b": 2,
	}
	assert.Equal(t, result, expectedResult)
}
