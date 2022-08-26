package lc_1971

func validPath(n int, edges [][]int, source int, destination int) bool {
	exist := make([]int, 0)
	return handle(&exist, edges, source, destination)
}

func handle(exist *[]int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	direct := checkHasDirectPath(edges, source, destination)
	if direct {
		return true
	}

	before := make([]int, 0)

	for _, item := range edges {
		if checkInArray(*exist, item[0]) || checkInArray(*exist, item[1]) {
			continue
		}
		if item[1] == destination {
			before = append(before, item[0])
		}
		if item[0] == destination {
			before = append(before, item[1])
		}
	}

	if len(before) == 0 {
		return false
	}

	*exist = append(*exist, destination)

	checkAll := make([]bool, len(before))

	for index, item := range before {
		checkAll[index] = handle(exist, edges, source, item)
	}

	return checkHasTrue(checkAll)
}

func checkHasDirectPath(edges [][]int, source int, destination int) bool {
	//check := make(map[string]bool)
	//for _, item := range edges {
	//	check[fmt.Sprintf("%d-%d", item[0], item[1])] = true
	//	check[fmt.Sprintf("%d-%d", item[1], item[0])] = true
	//}
	//
	//_, ok := check[fmt.Sprintf("%d-%d", source, destination)]
	//return ok
	for _, item := range edges {
		if item[0] == source && item[1] == destination {
			return true
		}
		if item[1] == source && item[0] == destination {
			return true
		}
	}
	return false
}

func checkHasTrue(array []bool) bool {
	for _, item := range array {
		if item {
			return true
		}
	}
	return false
}

func checkInArray(array []int, value int) bool {
	for _, item := range array {
		if value == item {
			return true
		}
	}
	return false
}
