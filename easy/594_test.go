package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findLHS(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{"for the example 1 from LeetCode", []int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{"for the variant example 1 from LeetCode", []int{1, 3, 2, 2, 5, 5, 2, 3, 7}, 5},
		{"for the variant example 1 from LeetCode", []int{1, 3, 2, 2, 5, 3, 2, 3, 7}, 6},
		{"for the test case from LeetCode", []int{1, 3, 5, 7, 9, 11, 13, 15, 17}, 0},
		{"for the test case from LeetCode", []int{1, 1, 1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLHS(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
