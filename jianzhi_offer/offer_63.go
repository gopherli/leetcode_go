package jianzhi_offer

// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/

func MaxProfit(prices []int) int {
	// dp[i] = max(dp[i-1],p[i]-min(p[0:i])) [1<=i]
	// i = 1 => max(dp[0],p[1]-p[0:1])
	// i = 2 => max(dp[1],p[2]-p[0:2])
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0] + 1
	max := 0

	for i := range prices {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
	}

	return max
}
