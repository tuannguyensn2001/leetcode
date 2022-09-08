package lc_94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	result := make([]int, 0)

	for _, item := range left {
		result = append(result, item)
	}

	result = append(result, root.Val)

	for _, item := range right {
		result = append(result, item)
	}

	return result

}
