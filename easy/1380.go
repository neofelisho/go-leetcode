package easy

func luckyNumbers(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix[0]) == 0 {
		return []int{}
	}
	currentLuckyNumbers := make(map[int]*int, len(matrix[0]))
	maxRowNumbers := make(map[int]*int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		var minPosition int
		var minValue int
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				minPosition = j
				minValue = matrix[i][j]

			} else if matrix[i][j] < minValue {
				minPosition = j
				minValue = matrix[i][j]
			}
			if currentLuckyNumbers[j] != nil && *currentLuckyNumbers[j] < matrix[i][j] {
				currentLuckyNumbers[j] = nil
			}
			if maxRowNumbers[j] == nil || *maxRowNumbers[j] < matrix[i][j] {
				maxRowNumbers[j] = &matrix[i][j]
			}
		}
		if currentLuckyNumbers[minPosition] == nil && minValue >= *maxRowNumbers[minPosition] {
			currentLuckyNumbers[minPosition] = &minValue
		}
	}
	results := make([]int, 0)
	for _, number := range currentLuckyNumbers {
		if number != nil {
			results = append(results, *number)
		}
	}
	return results
}
