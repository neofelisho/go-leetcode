package medium

func checkSubarraySum(nums []int, k int) bool {
	sumMap := make(map[int]int, len(nums)-1)
	sumMap[0] = -1
	sum := 0
	for i, num := range nums {
		sum += num
		if k != 0 {
			sum = sum % k
		}
		if valueOfIndex, ok := sumMap[sum]; ok {
			if i-valueOfIndex > 1 {
				return true
			}
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return false
}
