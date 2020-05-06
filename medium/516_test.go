package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"for the example 1 from LeetCode", "bbbab", 4},
		{"for the example 2 from LeetCode", "cbbd", 2},
		{"for the test case from LeetCode", "abaabaa", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindromeSubseq(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
