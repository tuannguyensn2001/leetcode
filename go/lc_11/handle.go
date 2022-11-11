package lc_11

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := min(height[left], height[right]) * (right - left)

	for {
		if left >= right {
			break
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		result = max(min(height[left], height[right])*(right-left), result)
	}

	return result

}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
