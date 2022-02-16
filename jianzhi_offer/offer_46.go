package jianzhi_offer

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

func TranslateNum(num int) int {
	// int转string
	s := strconv.Itoa(num)

	// p保留q,q保留r,r初始1（单数字翻译，为第一种情况）
	p, q, r := 0, 0, 1

	// 遍历，计算连续双数字翻译
	for i := 0; i < len(s); i++ {
		p, q, r = q, r, 0
		// r加上保留的上一次值
		r += q
		// 单字母翻译，跳过继续下一次
		if i == 0 {
			continue
		}
		// 从i=0开始，连续两个数字
		pre := s[i-1 : i+1]
		fmt.Println("pre", pre)
		// 满足条件即说明存在一个
		// 12344
		if pre >= "10" && pre <= "25" {
			r += p
		}
	}
	return r
}
