package medium

import "reflect"

func findAnagrams(s string, p string) []int {
	var results []int
	if len(p) > len(s) {
		return results
	}

	targetMap := map[rune]int{}
	for _, r := range p {
		targetMap[r]++
	}

	stateMap := map[rune]int{}
	for i, r := range s {
		stateMap[r]++
		if i < len(p) {
			if i == len(p)-1 && reflect.DeepEqual(targetMap, stateMap) {
				results = append(results, i-len(p)+1)
			}
		} else {
			removeKey := rune(s[i-len(p)])
			if count, ok := stateMap[removeKey]; ok {
				if count == 1 {
					delete(stateMap, removeKey)
				} else {
					stateMap[removeKey]--
				}
			}
			if reflect.DeepEqual(targetMap, stateMap) {
				results = append(results, i-len(p)+1)
			}
		}
	}

	return results
}
