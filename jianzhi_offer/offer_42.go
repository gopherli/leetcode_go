package jianzhi_offer

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

func MaxSubArray(nums []int) int {
	// 设第一个元素为最大连续子数组的和
	max := nums[0]

	// 从第二个元素开始计算i位置最大连续子数组的和
	for i := 1; i < len(nums); i++ {
		// 得出i位置的最大连续子数组的和
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 比较i位置最大连续子数组的和与之前最大max
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
