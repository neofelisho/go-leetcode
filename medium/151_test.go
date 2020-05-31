package medium

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"case 1", "the sky is blue", "blue is sky the"},
		{"case 2", "  hello world!  ", "world! hello"},
		{"case 3", "a good   example", "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
