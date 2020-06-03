package medium

func nextGreaterElements(nums []int) []int {
	var tmpNums []int
	tmpNums = append(tmpNums, nums...)
	tmpNums = append(tmpNums, nums...)

	results := make([]int, len(nums))
	for i := range results {
		results[i] = -1
	}

	for i, num := range nums {
		for j := i + 1; j < i+len(nums); j++ {
			if tmpNums[j] > num {
				results[i] = tmpNums[j]
				break
			}
		}
	}

	return results
}
