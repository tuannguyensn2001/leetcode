package lc_2011

func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, item := range operations {
		if item == "++X" || item == "X++" {
			x++
		} else if item == "--X" || item == "X--" {
			x--
		}
	}

	return x
}
