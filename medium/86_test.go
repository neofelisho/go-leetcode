package medium

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		list []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{"case 1", args{[]int{1, 4, 3, 2, 5, 2}, 3}, []int{1, 2, 2, 4, 3, 5}},
		//{"case 2", args{[]int{}, 3}, []int{}},
		//{"case 3", args{[]int{1}, 3}, []int{1}},
		//{"case 4", args{[]int{1, 1}, 0}, []int{1, 1}},
		{"case 5", args{[]int{2, 1}, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := GetListHead(tt.args.list)
			want := GetListHead(tt.want)
			if got := partition(head, tt.args.x); !reflect.DeepEqual(got, want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
