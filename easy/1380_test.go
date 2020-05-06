package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"for default case from LeetCode",
			args{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}},
			[]int{15},
		},
		{
			"for empty matrix",
			args{[][]int{}},
			[]int{},
		},
		{
			"for empty column",
			args{[][]int{{}, {}, {}}},
			[]int{},
		},
		{
			"for matrix with negative values",
			args{[][]int{{-3, 7, 8}, {-9, 11, 13}, {-15, 16, 17}}},
			[]int{-3},
		},
		{
			"for case from LeetCode",
			args{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}},
			[]int{12},
		},
		{
			"for variant case from LeetCode",
			args{[][]int{{1, 10, 4, 2}, {9, 3, 8, 13}, {15, 16, 17, 12}}},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := luckyNumbers(tt.args.matrix)
			assert.Equal(t, tt.want, got)
		})
	}
}
