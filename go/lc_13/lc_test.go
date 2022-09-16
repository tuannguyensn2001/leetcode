package lc_13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	input := "III"
	assert.Equal(t, 3, romanToInt(input))
}
