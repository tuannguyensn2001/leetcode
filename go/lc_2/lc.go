package lc_2

import (
	"math"
)

func sum(a []int, b []int) []int {
	result := make([]int, 0)

	diff := int(math.Abs(float64(len(a) - len(b))))

	for i := 1; i <= diff; i++ {
		if len(a) > len(b) {
			b = append(b, 0)
		} else {
			a = append(a, 0)
		}
	}

	remember := 0

	for i := 0; i < len(a); i++ {
		c := a[i] + b[i] + remember
		div := c % 10

		result = append(result, div)
		remember = c / 10

	}

	if remember != 0 {
		result = append(result, remember)
	}

	return result

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := convertNodeToArray(l1)
	b := convertNodeToArray(l2)

	c := sum(a, b)

	result := convertArrayToNode(c)
	return result
}

func convertNodeToArray(node *ListNode) []int {
	var result []int

	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

func convertArrayToNode(arr []int) *ListNode {

	node := &ListNode{
		Val: arr[0],
	}

	for i := 1; i < len(arr); i++ {
		push(node, arr[i])
	}

	return node

}

func push(node *ListNode, val int) {
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &ListNode{Val: val}
}
