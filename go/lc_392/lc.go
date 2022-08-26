package lc_392

func isSubsequence(s string, t string) bool {

	index := 0
	count := 0
	for _, item := range s {
		for i := index; i < len(t); i++ {
			if string(t[i]) == string(item) {
				index = i + 1
				count++
				break
			}
		}
	}

	return count == len(s)
}
