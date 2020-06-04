package medium

import (
	"sort"
)

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)

	// 2, 2, 3, 4
	// ^  ^     ^
	// i  j     k
	// nums[i] + nums[j] > nums[k]
	// when we move to next j, because next_j must >= prev_j, the above equation must be true
	// so we don't need to move the k back and recalculate again
	counts := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] <= 0 {
			continue
		}
		k := i + 2
		for j := i + 1; j < len(nums)-1; j++ {
			for k < len(nums) && nums[i]+nums[j] > nums[k] {
				k++
			}
			counts += k - j - 1
		}
	}
	return counts
}
