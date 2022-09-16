package lc_13

func romanToInt(s string) int {
	mapChar := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0

	sum += mapChar[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		next := mapChar[string(s[i+1])]
		current := mapChar[string(s[i])]

		if next > current {
			sum -= current
		} else {
			sum += current
		}

	}

	return sum

}
