package lcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"abdacbab", "acebfca", 4},
		{"", "acebfca", 0},
		{"abdacbab", "", 0},
		{"", "", 0},
		{"abcdef", "abcdef", 6},
		{"abc", "def", 0},
		{"AGGTAB", "GXTXAYB", 4},
		{"aaaa", "aa", 2},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, LCS(tt.s1, tt.s2))
	}
}
