package jianzhi_offer

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/

func MaxValue(grid [][]int) int {
	// 行，列
	m := len(grid)
	n := len(grid[0])

	// 初始化第一行
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	// 初始化第一列
	for j := 1; j < m; j++ {
		grid[j][0] += grid[j-1][0]
	}

	// 求最大
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += MaxNum(grid[i][j-1], grid[i-1][j])

		}
	}

	return grid[m-1][n-1]
}

func MaxNum(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
