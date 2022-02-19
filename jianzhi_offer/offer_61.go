package jianzhi_offer

// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

func IsStraight(nums []int) bool {
	// 排序
	quickSortIsStraight(nums, 0, len(nums)-1)

	// 1,2,3,4,5
	// 0,1,2,3,4
	// 0,1,3,4,5
	// 0,1,3,5,6
	// 0,0,1,3,5

	// 0的个数
	zero := 0
	for _, v := range nums {
		if v == 0 {
			zero++
		}
	}

	// 遍历
	nums = nums[zero:]
	for i := 1; i < len(nums); i++ {
		cha := nums[i] - nums[i-1]
		if cha == 1 {
			continue
		} else if cha > 1 {
			cha--
			zero = zero - cha
			if zero < 0 {
				return false
			}
			continue
		} else {
			return false
		}
	}

	return true
}

func quickSortIsStraight(arr []int, left, right int) {
	if left < right {
		privot := left
		index := left + 1
		for i := index; i <= right; i++ {
			if arr[privot] > arr[i] {
				arr[index], arr[i] = arr[i], arr[index]
				index++
			}
		}
		arr[privot], arr[index-1] = arr[index-1], arr[privot]
		partition := index - 1
		quickSortIsStraight(arr, left, partition-1)
		quickSortIsStraight(arr, partition+1, right)
	}
}
