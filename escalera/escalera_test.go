package escalera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEscalera(t *testing.T) {
	tests := []struct {
		n        int
		formas   []int
		expected int
	}{
		{10, []int{2, 4, 5, 8}, 11},
		{0, []int{1, 2}, 1}, // caso base: cero escalones
		{1, []int{2, 3}, 0}, // sin forma de subir 1 escalón con pasos de 2 o 3
		{5, []int{1}, 1},    // solo movimientos de 1 escalón
		{4, []int{1, 2}, 5}, // clásico de 1 o 2 escalones
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, Escalera(tt.n, tt.formas))
	}
}
