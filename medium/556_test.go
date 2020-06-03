package medium

import "testing"

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{"case 1", 12, 21},
		{"case 2", 21, -1},
		{"case 3", 12345, 12354},
		{"case 4", 12354, 12435},
		{"case 5", 230241, 230412},
		{"case 6", 12443322, 13222344},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
