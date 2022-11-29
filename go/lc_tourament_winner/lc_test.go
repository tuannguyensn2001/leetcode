package lc_tourament_winner

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestTournamentWinner(t *testing.T) {
	tests := []struct {
		Competitions [][]string
		Result       []int
		Output       string
	}{
		{
			Competitions: [][]string{
				{"HTML", "C#"},
				{"C#", "Python"},
				{"Python", "HTML"},
			},
			Result: []int{0, 0, 1},
			Output: "Python",
		},
	}

	for index, item := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			output := TournamentWinner(item.Competitions, item.Result)
			assert.Equal(t, item.Output, output)
		})
	}
}
