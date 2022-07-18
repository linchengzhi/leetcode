package _00_899

func superEggDrop(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for m := 1; m <= N; m++ {
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
			if dp[k][m] >= N {
				return m
			}
		}
	}
	return N
}
