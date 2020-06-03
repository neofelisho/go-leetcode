package medium

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1) + 1
	m := len(text2) + 1
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = getMax(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

func getMax(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
