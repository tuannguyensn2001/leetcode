package lc_1025

func divisorGame(n int) bool {
	a := getDivide(n)

	for _, item := range a {
		b := n - item
		if !handle(b, 1) {
			return false
		}
	}

	return true

}

func handle(n int, turn int) bool {
	if turn%2 == 1 {
		turn = 1
	} else if turn%2 == 0 {
		turn = 2
	}
	if n == 1 && turn == 2 {
		return false
	}

	if n == 1 && turn == 1 {
		return true
	}

	a := getDivide(n)

	check := make([]bool, len(a))

	for index, item := range a {
		check[index] = handle(n-item, turn+1)
	}

	for _, item := range check {
		if !item {
			return false
		}
	}

	return true
}

func getDivide(n int) []int {
	result := make([]int, 0)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}
