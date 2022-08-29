package lc_2016

func maximumDifference(nums []int) int {
	result := -1
	minValue := nums[0]

	for i := 1; i < len(nums); i++ {
		if minValue > nums[i] {
			minValue = nums[i]
		} else if nums[i]-minValue > result {
			result = nums[i] - minValue
		}
	}

	if result == 0 {
		return -1
	}

	return result
}
