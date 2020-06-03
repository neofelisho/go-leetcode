package medium

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"case 1", []int{1, 2, 1}, []int{2, -1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
