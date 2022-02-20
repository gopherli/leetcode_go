package jianzhi_offer

// https://leetcode-cn.com/problems/chou-shu-lcof/

func NthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = Min(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}

	return dp[n-1]
}

func Min(a, b, c int) int {
	min := 0
	if a > b {
		min = b
		if b > c {
			min = c
		}
	} else {
		min = a
		if a > c {
			min = c
		}
	}

	return min
}
