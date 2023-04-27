package subconjuntos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumaSubconjunto(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4, 3},
		{[]int{1, 2, 3, 1, 4}, 6, 4},
		{[]int{2, 4, 2, 6, 8}, 7, 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, SumaSubconjunto(tt.arr, tt.k))
	}
}
