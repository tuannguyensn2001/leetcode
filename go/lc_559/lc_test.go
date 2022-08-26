package lc_559

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxDepth(t *testing.T) {

	node1 := Node{
		Val:      1,
		Children: make([]*Node, 0),
	}
	node2 := Node{
		Val:      1,
		Children: make([]*Node, 0),
	}
	node3 := Node{
		Val:      1,
		Children: make([]*Node, 0),
	}

	node := Node{
		Val:      1,
		Children: []*Node{&node1, &node2, &node3},
	}

	output := maxDepth(&node)

	assert.Equal(t, 2, output)

}
