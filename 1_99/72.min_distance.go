package __99

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 && n == 0 {
		return 0
	}
	var dp = make([][]int, 0)
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min3(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}

	return dp[m][n]
}

func min3(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}
