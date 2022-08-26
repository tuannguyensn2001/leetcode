package lc_2011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFinal(t *testing.T) {
	input := []string{"++X", "++X", "X++"}

	output := finalValueAfterOperations(input)

	assert.Equal(t, 3, output)

}
