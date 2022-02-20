package jianzhi_offer

import "strconv"

// https://leetcode-cn.com/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/

func FindNthDigit(n int) int {
	// 	dight  start  nums  count
	//  1		1		9		9
	//  2		10		90		180
	//	3		100		900		2700

	// 获取位数，开始数
	dight, start, count := 1, 1, 9
	for n > count {
		n -= count
		dight += 1
		start *= 10
		count = dight * start * 9
	}

	// 获取所在数字
	num := start + (n-1)/2

	// 获取位数所在数字的数
	i := (n - 1) % dight
	s := strconv.Itoa(num)
	b := s[i]
	number, _ := strconv.Atoi(string(b))

	return number
}
