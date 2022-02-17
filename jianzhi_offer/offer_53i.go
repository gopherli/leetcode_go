package jianzhi_offer

// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/

func Search(nums []int, target int) int {
	count := 0
	for _, v := range nums {
		if v > target {
			break
		}

		if v == target {
			count++
		}
	}
	return count
}
