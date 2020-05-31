package medium

// Good for interview question
// Key point is to keep the previous sum results
func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int, len(nums)-1)
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		if sum == k {
			count++
		}
		if valueCount, ok := sumMap[sum-k]; ok {
			count += valueCount
		}
		sumMap[sum] += 1
	}
	return count
}
