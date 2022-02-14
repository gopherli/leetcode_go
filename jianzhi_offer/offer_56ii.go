package jianzhi_offer

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/

func SingleNumber56II(nums []int) int {
	return QuickSortFor56II(nums)
}

func QuickSortFor56II(nums []int) int {
	// 快排
	quickSort56I(nums, 0, len(nums)-1)

	// 头处理
	if nums[0] < nums[1] {
		return nums[0]
	}

	// 尾处理
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return nums[len(nums)-1]
	}

	// 遍历
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			return nums[i]
		}
	}

	return 0
}
