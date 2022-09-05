package lc_429

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	t.Run("case-1", func(t *testing.T) {
		root := Node{
			Val: 1,
		}

		output := levelOrder(&root)

		expected := [][]int{{1}}
		assert.Equal(t, expected, output)

	})

	t.Run("case-2", func(t *testing.T) {
		root := &Node{
			Val: 1,
			Children: []*Node{
				&Node{
					Val: 2,
					Children: []*Node{
						&Node{
							Val: 5,
						},
					},
				},
				&Node{
					Val: 3,
				},
				&Node{
					Val: 4,
				},
			},
		}

		output := levelOrder(root)

		expected := [][]int{{1}, {2, 3, 4}, {5}}

		assert.Equal(t, expected, output)
	})
}
