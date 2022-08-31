package lc_417

import "testing"

func TestPacificAtlantic(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		input := [][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}

		pacificAtlantic(input)

	})
}
