package main

import "fmt"

func main() {
	fmt.Println(minimumDeletions("ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa"))
}

func minimumDeletions(s string) int {
	t, l := 0, len(s)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		if s[l-i-1] == 'a' {
			t++
		}
		dp[l-1-i] = t
	}
	t = 0
	n, ans := 0, 100001
	for i := 0; i < l; i++ {
		if s[i] == 'b' {
			if t == 0 {
				ans = min(ans, n+dp[i])
			}
			t++
		} else if s[i] == 'a' {
			n += t
			t = 0
		}
	}
	return min(n, ans)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
