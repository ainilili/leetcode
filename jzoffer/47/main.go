package main

import "fmt"

func main() {
	//fmt.Println(maxValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(maxValue([][]int{{1, 2, 5}, {3, 2, 1}}))
}

func maxValue(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i, arr := range grid {
		dp[i] = make([]int, len(arr))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = max(dp[i][j-1]+grid[i][j], dp[i-1][j]+grid[i][j])
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
