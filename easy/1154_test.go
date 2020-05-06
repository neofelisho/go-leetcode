package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dayOfYear(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"for the example 1 from LeetCode", "2019-01-09", 9},
		{"for the example 2 from LeetCode", "2019-02-10", 41},
		{"for the example 3 from LeetCode", "2003-03-01", 60},
		{"for the example 4 from LeetCode", "2004-03-01", 61},
	}
	for _, tt := range tests {
		got := dayOfYear(tt.arg)
		assert.Equal(t, tt.want, got)
	}
}
