package mochila

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMochila(t *testing.T) {
	tests := []struct {
		valores  []int
		pesos    []int
		k        int
		expected int
	}{
		{[]int{20, 30, 15, 25, 10}, []int{6, 13, 5, 10, 3}, 20, 55},
		{[]int{5, 10, 15}, []int{1, 3, 4}, 0, 0},
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 220},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, Mochila(tt.valores, tt.pesos, tt.k))
	}
}
