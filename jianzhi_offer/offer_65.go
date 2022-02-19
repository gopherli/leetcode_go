package jianzhi_offer

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/

// 位运算：加减法
func Add(a int, b int) int {
	// 加法：1. 未进位和 2. 进位
	// 位运算加法： 1. + 2.
	// 1. a^b + 2.a&b<<1
	// n := a ^ b
	// c := a & b << 1
	for b != 0 {
		c := a & b << 1 // 进位
		a ^= b          // 	未进位和（b第一次没带进位，b第二次带进位）
		b = c           // 进位
	}
	return a
}

// 位运算：乘法
func Multi(a, b int) int {
	return AddFor(a, b)
}

func AddFor(a int, count int) int {
	// c := a & b << 1
	a1 := a
	for count > 1 {
		b := a1
		for b != 0 {
			c := a & b << 1 // 进位
			a ^= b          // 	未进位和（b第一次没带进位，b第二次带进位）
			b = c           // 进位
		}
		count--
	}

	return a
}
