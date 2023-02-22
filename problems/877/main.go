package main

import "math"

func stoneGameII(piles []int) int {
	dp := [101][101]int{}
	dp[0][1] = piles[0]
	dp[0][2] = piles[0] + piles[1]
	l := len(piles)
	ans := 0
	for i := 2; i < l; i++ {
		for j := i + 1; j < min(i+sqrt(i), l); j++ {
			dp[i][j] = max(dp[i-2][j-1]-piles[j], dp[i-2][j-2]-piles[j])
			ans = max(dp[i][j], ans)
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}
