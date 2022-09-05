package lc_429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		length := len(queue)
		temp := make([]int, 0)
		for _, item := range queue {
			temp = append(temp, item.Val)

			for _, child := range item.Children {
				if child != nil {
					queue = append(queue, child)
				}
			}
		}

		result = append(result, temp)
		if len(queue) == length {
			break
		}
		queue = queue[length:]
	}

	return result
}
