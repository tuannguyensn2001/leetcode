package lc_tourament_winner

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	m := make(map[string]int)

	for index, item := range competitions {
		var win string
		if results[index] == 1 {
			win = item[0]
		} else {
			win = item[1]
		}
		m[win] += 3
	}

	return getKeyMax(m)
}

func getKeyMax(m map[string]int) string {
	max := -1
	result := ""
	for key, val := range m {
		if max < val {
			max = val
			result = key
		}

	}
	return result
}
