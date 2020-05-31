package medium

import "strings"

func reverseWords(s string) string {
	splits := strings.Split(s, " ")
	var result string
	for i := len(splits) - 1; i >= 0; i-- {
		split := splits[i]
		if split == "" {
			continue
		}
		if result == "" {
			result += split
		} else {
			result += " " + split
		}
	}
	return result
}
