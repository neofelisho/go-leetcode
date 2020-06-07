package medium

func longestMountain(A []int) int {
	maxHeight := 0
	start := -1
	isUphill := false
	for i := 1; i < len(A); i++ {
		if A[i-1] < A[i] {
			// Uphill
			if !isUphill && start != -1 {
				// End of the mountain
				height := i - start
				maxHeight = getMax(maxHeight, height)
				start = -1
			}
			if start == -1 {
				// New mountain
				start = i - 1
				isUphill = true
			}
		} else if A[i-1] > A[i] {
			// Downhill
			isUphill = false
		} else {
			if !isUphill && start != -1 {
				// End of the mountain
				height := i - start
				maxHeight = getMax(maxHeight, height)
			}
			//	Reset
			start = -1
			isUphill = false
		}
	}
	if start != -1 && !isUphill {
		maxHeight = getMax(maxHeight, len(A)-start)
	}
	return maxHeight
}
