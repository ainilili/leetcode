package _ms1_9mb

import "fmt"

func main() {
	fmt.Println(climbStairs(44))
}

var dp = make([]int, 45)
var r = 0

func climbStairs(n int) int {
	return fn(int16(n), 1) + fn(int16(n), 2)
}

func fn(n int16, d int16) int{
	n = n - d
	if n == 0{
		return 1
	}
	if n < 0{
		return 0
	}
	if r = dp[n]; r > 0{
		return r
	}else{
		dp[n] = fn(n, 1) + fn(n, 2)
		return dp[n]
	}
}