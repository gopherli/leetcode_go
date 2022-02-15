package jianzhi_offer

// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

func FindContinuousSequence(target int) [][]int {
	// 1.枚举法
	// return Enumerate(target)

	// 2.双指针
	return DoublePointer(target)
}

// 1. 枚举法
func Enumerate(target int) [][]int {
	// 满足条件的结果集合
	res := make([][]int, 0)

	// 枚举次数
	n := target / 2

	// 执行枚举
	for i := 1; i <= n; i++ {
		// 开始
		sum := 0
		// 从哪个数开始枚举
		for j := i; ; j++ {
			sum += j
			if sum > target {
				break
			} else if sum == target {
				// 满足条件的数组集合
				rec := make([]int, 0)
				for k := i; k <= j; k++ {
					rec = append(rec, k)
				}
				res = append(res, rec)
				break
			}
		}
	}

	return res
}

// 2. 双指针
func DoublePointer(target int) [][]int {
	// 结果
	res := make([][]int, 0)

	// 双指针
	l, r := 1, 2

	// 滑动窗口,l,r只增不减，某一时刻只有一个指针滑动
	for l < r {
		// l到r的正整数求和公式:(首项+末项)*项数[末项-首项+1]/2
		sum := (l + r) * (r - l + 1) / 2
		// 满足
		if sum == target {
			one := make([]int, 0)
			for i := l; i <= r; i++ {
				one = append(one, i)
			}
			res = append(res, one)
			l++
			// 大于
		} else if sum > target {
			l++
			// 小于
		} else {
			r++
		}
	}

	return res
}
