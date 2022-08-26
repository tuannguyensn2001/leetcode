package lc_1971

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidPath(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)
		assert.True(t, output)
	})

	t.Run("case 2", func(t *testing.T) {
		output := validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}, {2, 3}}, 0, 5)
		assert.True(t, output)
	})
}
