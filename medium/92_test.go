package medium

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type arg struct {
		list []int
		m    int
		n    int
	}
	tests := []struct {
		name string
		args arg
		want []int
	}{
		{"case 1", arg{[]int{1, 2, 3, 4, 5}, 2, 4}, []int{1, 4, 3, 2, 5}},
		{"case 2", arg{[]int{5}, 1, 1}, []int{5}},
		{"case 3", arg{[]int{3, 5}, 1, 1}, []int{3, 5}},
		{"case 4", arg{[]int{3, 5}, 1, 2}, []int{5, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := GetListHead(tt.args.list)
			expected := GetListHead(tt.want)
			if got := reverseBetween(head, tt.args.m, tt.args.n); !reflect.DeepEqual(got, expected) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
