package easy

func findLHS(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	currentLHS := 0
	for key, value := range counter {
		if val, ok := counter[key+1]; ok {
			currentLHS = max(currentLHS, value+val)
		}
	}

	return currentLHS
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
