package lc_1504

type Point struct {
	x int
	y int
}

func getInRow(mat [][]int, x int, y int) int {
	count := 0

	for i := y + 1; i < len(mat[0]); i++ {
		if mat[x][i] == 1 && mat[x][i-1] == 1 {
			count++
		}
	}

	return count
}

func getInColumn(mat [][]int, x int, y int) int {
	count := 0

	for i := x + 1; i < len(mat); i++ {
		if mat[i][y] == 1 && mat[i][y-1] == 1 {
			count++
		}
	}

	return count
}
