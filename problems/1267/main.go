package main

import "fmt"

func main() {
	fmt.Println(countServers([][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}))
}
func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mp, np := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[i] += grid[i][j]
			np[j] += grid[i][j]
		}
	}
	ans := 0
	for _, v := range mp {
		if v > 1 {
			ans += v
		}
	}
	for j, v := range np {
		if v < 2 {
			continue
		}
		for i := 0; i < m; i++ {
			if mp[i] > 1 && grid[i][j] == 1 {
				v--
			}
		}
		ans += v
	}
	return ans
}
