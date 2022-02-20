package jianzhi_offer

import "math"

// https://leetcode-cn.com/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/

func StrToInt(str string) int {
	// 去除头部空格
	// 字符转数字：num := int(n-'0')
	// 数字拼接：res := 0 res = res*10+num

	// 1.去除头部空格
	b := []byte(str)
	left, right := 0, len(b)
	for left < right && b[left] == ' ' {
		left++
	}
	b = b[left:]

	// 结果
	res := 0
	// 符号
	sign := 1

	// 操作
	for i, v := range b {
		if v >= '0' && v <= '9' {
			// 数字拼接 res = res * 10+n
			// 字符转数字 int(v-'0')
			res = res*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}
		// 大数操作
		if res > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
	}

	return sign * res
}
