package lc_559

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	maxList := make([]int, 0)

	for _, item := range root.Children {
		maxList = append(maxList, maxDepth(item))
	}

	max := maxList[0]

	for i := 1; i < len(maxList); i++ {
		if max < maxList[i] {
			max = maxList[i]
		}
	}

	return max + 1

}
