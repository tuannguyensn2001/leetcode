package lc_112

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	t.Run("leaf", func(t *testing.T) {
		node := &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}

		output := hasPathSum(node, 1)

		assert.True(t, output)
	})

	t.Run("2 leaf", func(t *testing.T) {
		left := &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}
		right := &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		}

		root := &TreeNode{
			Val:   3,
			Left:  left,
			Right: right,
		}

		output := hasPathSum(root, 4)

		assert.True(t, output)

		output = hasPathSum(root, 5)
		assert.True(t, output)

	})
}
