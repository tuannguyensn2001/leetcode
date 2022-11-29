package lc_977

func sortedSquares(array []int) []int {
	result := make([]int, len(array))
	left := 0
	right := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		a := array[left]
		b := array[right]
		if abs(b) <= abs(a) {
			result[i] = a * a
			left += 1
		} else {
			result[i] = b * b
			right -= 1
		}
	}

	return result

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
