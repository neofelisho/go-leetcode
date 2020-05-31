package medium

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1", args{nums: []int{1, 1, 1}, k: 2}, 2},
		{"case 2", args{nums: []int{1, 2, 1, 2, 1}, k: 3}, 4},
		{"case 3", args{nums: []int{28, 54, 7, -70, 22, 65, -6}, k: 100}, 1},
		{"case 4", args{nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, k: 0}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
