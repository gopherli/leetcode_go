package jianzhi_offer

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

func MissingNumber(nums []int) int {
	// 	不重复排序数组，优先二分法
	left, right := 0, len(nums)-1
	// 一定要<=
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
