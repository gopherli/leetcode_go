package jianzhi_offer

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

func MissingNumber(nums []int) int {
	misNum := 0

	if len(nums) == 1 {
		if nums[0] > 0 {
			misNum = nums[0] - 1
		} else {
			misNum = nums[0] + 1
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > 1 {
			misNum = nums[i] - 1
			break
		}
	}
	return misNum
}
