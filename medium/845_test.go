package medium

import "testing"

func Test_longestMountain(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"case 1", []int{2, 1, 4, 7, 3, 2, 5}, 5},
		{"case 2", []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}, 11},
		{"case 3", []int{0, 2, 2}, 0},
		{"case 4", []int{875, 884, 239, 731, 723, 685}, 4},
		{"case 5", []int{2, 2, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
