package jianzhi_offer

// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/

func ConstructArr(a []int) []int {
	//  6 2 3 4 5
	//0 1 2 3 4 5
	//1 6 1 3 4 5
	//2 6 2 1 4 5
	//3 6 2 3 1 5
	//4 6 2 3 4 1
	// 上三角
	// dp[i] = dp[i-1]*a[i-1]
	// 下三角
	// dp[i] = tmp*a[i+1]
	res := make([]int, len(a))
	res[0] = 1 // 上三角：第一个元素初始化为1
	tmp := 1   // 下三角：最后一个元素初始化为1

	// 上三角
	for i := 1; i < len(a); i++ {
		res[i] = res[i-1] * a[i-1]
	}

	// 下三角
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		res[i] *= tmp
	}

	return res
}
