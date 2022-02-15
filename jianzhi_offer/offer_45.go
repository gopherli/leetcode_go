package jianzhi_offer

import "strconv"

// https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/

func MinNumber(arr []int) string {
	// 字符串数组
	strs := make([]string, 0)
	for i := range arr {
		strs = append(strs, strconv.Itoa(arr[i]))
	}

	// 快排,得出最小的值
	_quickSortMinNumber(strs, 0, len(arr)-1)

	// 字符串数组转字符串
	s := ""
	for i := range strs {
		s += strs[i]
	}

	return s
}

// 快排
func _quickSortMinNumber(arr []string, left, right int) {
	if left < right {
		privot := left
		index := left + 1
		for i := index; i <= right; i++ {
			// 修改比较条件，即可
			if arr[privot]+arr[i] > arr[i]+arr[privot] {
				arr[i], arr[index] = arr[index], arr[i]
				index++
			}
		}
		arr[privot], arr[index-1] = arr[index-1], arr[privot]
		partitionIndex := index - 1
		_quickSortMinNumber(arr, left, partitionIndex-1)
		_quickSortMinNumber(arr, partitionIndex+1, right)
	}
}
