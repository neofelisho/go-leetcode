package medium

import "reflect"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	targetMap := map[rune]int{}
	for _, r := range s1 {
		targetMap[r]++
	}

	stateMap := map[rune]int{}
	windowSize := len(s1)
	for i, r := range s2 {
		if i < windowSize {
			stateMap[r]++
			if i == windowSize-1 && reflect.DeepEqual(stateMap, targetMap) {
				return true
			}
		} else {
			stateMap[r]++
			removeKey := rune(s2[i-windowSize])
			if value, ok := stateMap[removeKey]; ok {
				if value == 1 {
					delete(stateMap, removeKey)
				} else {
					stateMap[removeKey]--
				}
			}
			if reflect.DeepEqual(stateMap, targetMap) {
				return true
			}
		}
	}
	return false
}
