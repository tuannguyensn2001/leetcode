package lc_1025

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivisorGame(t *testing.T) {
	tests := []struct {
		Input  int
		Output bool
	}{
		{
			Input:  2,
			Output: true,
		},
		{
			Input:  3,
			Output: false,
		},
		{
			Input:  4,
			Output: true,
		},
		{
			Input:  5,
			Output: false,
		},
	}

	for index, item := range tests {
		t.Run(fmt.Sprintf("test-%v", index), func(t *testing.T) {
			output := divisorGame(item.Input)
			assert.Equal(t, item.Output, output)
		})
	}

}

func TestGetDivide(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		output := getDivide(4)
		assert.Equal(t, []int{1, 2}, output)
	})

	t.Run("case 2", func(t *testing.T) {
		output := getDivide(10)
		assert.Equal(t, []int{1, 2, 5}, output)
	})
}

func TestHandle(t *testing.T) {
	tests := []struct {
		N      int
		Turn   int
		Output bool
	}{
		{
			N:      1,
			Turn:   1,
			Output: true,
		},
		{
			N:      1,
			Turn:   2,
			Output: false,
		},
		{
			N:      4,
			Turn:   1,
			Output: true,
		},
	}

	for index, item := range tests {
		t.Run(fmt.Sprintf("case-%v", index), func(t *testing.T) {
			output := handle(item.N, item.Turn)
			assert.Equal(t, item.Output, output)
		})
	}

}
