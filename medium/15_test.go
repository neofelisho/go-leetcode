package medium

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{"case 1", []int{-4, -1, -1, 0, 1, 2}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"case 2", []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
