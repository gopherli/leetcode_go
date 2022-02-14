package jianzhi_offer

// https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/

func TwoSum(nums []int, target int) []int {
	// 数组nums是递增的,可以采用双指针

	// 头尾指针，两数之和
	s, i, j := 0, 0, len(nums)-1

	// 双指针操作
	for i < j {
		// 两数和
		s = nums[i] + nums[j]

		// 满足条件,返回结果
		if s == target {
			return []int{nums[i], nums[j]}
		}

		if s > target {
			// 尾指针减小，取更小数
			j--
		} else {
			// 头指针增大，取更大数
			i++
		}
	}

	return nil
}
