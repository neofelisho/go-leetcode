package medium

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{nums: []int{23, 2, 4, 6, 7}, k: 6}, true},
		{"case 2", args{nums: []int{23, 2, 6, 4, 7}, k: 6}, true},
		{"case 3", args{nums: []int{23, 2, 6, 4, 7}, k: 0}, false},
		{"case 4", args{nums: []int{0, 0}, k: 0}, true},
		{"case 5", args{nums: []int{1, 0}, k: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
