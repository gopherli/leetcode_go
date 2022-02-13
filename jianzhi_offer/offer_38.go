package jianzhi_offer

import (
	"sort"
)

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

func Permutation(s string) []string {
	// string转[]byte
	t := []byte(s)
	// 排序，防止排列出现重复元素
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	// 一个排列的空位数
	n := len(t)
	// 一个排列（如果给定len=0,append会在原有的基础上追加，导致"\b0000\b0000\b0000abs"）
	perm := make([]byte, 0)
	// 一个排列需要填充到位置,需要给定长度n
	vis := make([]bool, n)
	// 不重复全排列结果
	ans := make([]string, 0)

	// 定义回溯方法，每次return生成一个排列
	var backtracking func(int)
	// 回溯方法实现
	backtracking = func(i int) {
		// 生成一个排列
		if i == n {
			// []byte转string
			ans = append(ans, string(perm))
			// 结束当前回溯方法
			return
		}

		// 生成全部不重复排列
		for j, b := range vis {
			// 防止重复元素出现，当前元素true或者当前元素与上一个元素相同
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}

			// 元素参与排列，true
			vis[j] = true
			// 元素填充排列当前位置
			perm = append(perm, t[j])
			// 回溯，进行下一个位置的填充
			backtracking(i + 1)
			// 回溯完成，i=n完成一个排列，防止perm的append添加上一个结果
			perm = perm[:len(perm)-1]
			// 回溯完成，i=n完成一个排列，当前j复原false
			vis[j] = false
		}
	}
	// 调用回溯方法，第一个位置开始填充，i=0
	backtracking(0)

	return ans
}
