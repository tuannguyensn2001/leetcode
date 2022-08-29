package lc_1504

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInRow(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := [][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}

		output := getInRow(input, 0, 0)
		assert.Equal(t, 0, output)
	})

	t.Run("Case 2", func(t *testing.T) {
		input := [][]int{
			{1, 1, 0, 1},
			{1, 1, 0, 1},
			{1, 1, 0, 1},
		}

		output := getInRow(input, 0, 0)
		assert.Equal(t, 1, output)
	})
}

func TestGetInColumn(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		input := [][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}

		output := getInRow(input, 0, 0)
		assert.Equal(t, 2, output)
	})
}
