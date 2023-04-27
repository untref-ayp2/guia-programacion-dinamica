package caminominimo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCostoCaminoMinimo(t *testing.T) {
	tests := []struct {
		matriz   [][]int
		expected int
	}{
		// ejemplo original
		{[][]int{
			{3, 2, 12, 15, 10},
			{6, 19, 7, 11, 17},
			{8, 5, 12, 32, 21},
			{3, 20, 2, 9, 7},
		}, 52},
		// matriz 1x1
		{[][]int{{5}}, 5},
		// matriz 2x2
		{[][]int{
			{1, 2},
			{3, 4},
		}, 7},
		// ejemplo clásico 3x3
		{[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}, 7},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, CostoCaminoMinimo(tt.matriz))
	}
}
