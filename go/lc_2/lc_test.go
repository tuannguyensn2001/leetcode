package lc_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		a := []int{9, 9, 9, 9, 9, 9, 9}
		b := []int{9, 9, 9, 9}

		expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

		output := sum(a, b)

		assert.Equal(t, expected, output)
	})

	t.Run("case 2", func(t *testing.T) {
		a := []int{2, 4, 3}
		b := []int{5, 6, 4}

		expected := []int{7, 0, 8}

		output := sum(a, b)

		assert.Equal(t, expected, output)
	})
}

func TestConvertToArray(t *testing.T) {
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	expected := []int{1, 2, 3}
	output := convertNodeToArray(&node)

	assert.Equal(t, expected, output)
}

func TestConvertToNode(t *testing.T) {
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	output := convertArrayToNode([]int{1, 2, 3})

	assert.EqualValues(t, expected, output)
}

func TestAddTwoNumber(t *testing.T) {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	output := addTwoNumbers(&l1, &l2)

	assert.Nil(t, output)

}
