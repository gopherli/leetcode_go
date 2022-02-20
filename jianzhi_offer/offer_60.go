package jianzhi_offer

// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/

func DicesProbability(n int) []float64 {
	// 初始化骰子为1时的概率
	dp := make([]float64, 0)
	for i := 1; i <= 6; i++ {
		dp = append(dp, float64(1)/6)
	}

	// 几个骰子就循环几次
	for i := 2; i <= n; i++ {
		// 几个骰子对应的和的个数
		tmp := make([]float64, 5*i+1)
		// 上一个骰子的概率
		for j := 0; j < len(dp); j++ {
			// 当前骰子的概率
			for k := 0; k < 6; k++ {
				tmp[j+k] += dp[j] / float64(6)
			}
		}
		// 当前骰子概率
		dp = tmp
	}

	return dp
}
