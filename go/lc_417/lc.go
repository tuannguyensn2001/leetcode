package lc_417

import (
	"strconv"
)

func pacificAtlantic(heights [][]int) [][]int {
	xMax := len(heights)
	yMax := len(heights[0])
	var result [][]int

	pac := make(map[string]bool)
	alt := make(map[string]bool)

	for i := 0; i < yMax; i++ {
		dfs(0, i, pac, heights[0][i], xMax, yMax, heights)
		dfs(xMax-1, i, alt, heights[yMax-1][i], xMax, yMax, heights)
	}

	for i := 0; i < xMax; i++ {
		dfs(i, 0, pac, heights[i][0], xMax, yMax, heights)
		dfs(i, yMax-1, alt, heights[i][yMax-1], xMax, yMax, heights)
	}

	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			_, ok1 := pac[strconv.Itoa(i)+"-"+strconv.Itoa(j)]
			_, ok2 := alt[strconv.Itoa(i)+"-"+strconv.Itoa(j)]
			if ok1 && ok2 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfs(x int, y int, check map[string]bool, prevHeight int, xMax int, yMax int, heights [][]int) {
	_, ok := check[strconv.Itoa(x)+"-"+strconv.Itoa(y)]
	if ok || x < 0 || y < 0 || x >= xMax || y >= yMax || heights[x][y] < prevHeight {
		return
	}

	check[strconv.Itoa(x)+"-"+strconv.Itoa(y)] = true

	dfs(x-1, y, check, heights[x][y], xMax, yMax, heights)
	dfs(x, y-1, check, heights[x][y], xMax, yMax, heights)
	dfs(x+1, y, check, heights[x][y], xMax, yMax, heights)
	dfs(x, y+1, check, heights[x][y], xMax, yMax, heights)

}
