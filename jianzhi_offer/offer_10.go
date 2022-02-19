package jianzhi_offer

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

func NumWays(n int) int {
	dp0 := 1
	dp1 := 1
	for i := 0; i < n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp0
}
