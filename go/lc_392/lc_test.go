package lc_392

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		S      string
		T      string
		Output bool
	}{
		{
			S:      "abc",
			T:      "ahbgdc",
			Output: true,
		},
		{
			S:      "bb",
			T:      "ahbgdc",
			Output: false,
		},
	}

	for index, item := range tests {
		t.Run(fmt.Sprintf("case-%v", index), func(t *testing.T) {
			output := isSubsequence(item.S, item.T)
			assert.Equal(t, item.Output, output)
		})
	}

}
